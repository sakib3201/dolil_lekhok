package database

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"dolil_lekhok/backend/auth"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
)

var db *sql.DB

func InitDB() error {
	// Create data directory if it doesn't exist
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	// Open database with WAL mode
	var err error
	dbPath := filepath.Join(dataDir, "dolil_lekhok.db")
	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?_journal=WAL&_timeout=5000&_fk=true", dbPath))
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Enable foreign keys
	if _, err = db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	// Create users table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS user_images (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			original_filename TEXT NOT NULL,
			stored_filename TEXT NOT NULL UNIQUE,
			file_path TEXT NOT NULL,
			file_size INTEGER NOT NULL,
			mime_type TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	return nil
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Don't include password in JSON responses
}

func CreateUser(username, password string) (*User, error) {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	result, err := db.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		username, hashedPassword,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID: %v", err)
	}

	return &User{
		ID:       int(id),
		Username: username,
		Password: hashedPassword,
	}, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = ?",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func GetUserByID(id int) (*User, error) {
	var user User
	err := db.QueryRow(
		"SELECT id, username, password FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func AuthenticateUser(username, password string) (*User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil // User not found
	}

	err = auth.CheckPassword(user.Password, password)
	if err != nil {
		return nil, nil // Invalid password
	}

	// Don't return the hashed password
	user.Password = ""
	return user, nil
}

type ImageMetadata struct {
	ID              int
	UserID          int
	OriginalFilename string
	StoredFilename  string
	FilePath        string
	FileSize        int64
	MIMEType        string
	CreatedAt       string
}

// SaveImage saves the uploaded file and stores its metadata in the database
func SaveImage(userID int, fileHeader *multipart.FileHeader, uploadDir string) (*ImageMetadata, error) {
	// Ensure upload directory exists
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Open the uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Generate a unique filename
	ext := filepath.Ext(fileHeader.Filename)
	storedFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(uploadDir, storedFilename)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy the file content
	fileSize, err := io.Copy(dst, src)
	if err != nil {
		os.Remove(filePath) // Clean up if copy fails
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// Get MIME type
	file, err := os.Open(filePath)
	if err != nil {
		os.Remove(filePath) // Clean up if we can't determine MIME type
		return nil, fmt.Errorf("failed to open saved file: %v", err)
	}
	defer file.Close()

	// Read first 512 bytes to determine MIME type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		os.Remove(filePath) // Clean up if we can't read the file
		return nil, fmt.Errorf("failed to read file for MIME type: %v", err)
	}

	mimeType := http.DetectContentType(buffer)

	// Save metadata to database
	result, err := db.Exec(
		`INSERT INTO user_images 
		(user_id, original_filename, stored_filename, file_path, file_size, mime_type) 
		VALUES (?, ?, ?, ?, ?, ?)`,
		userID, fileHeader.Filename, storedFilename, filePath, fileSize, mimeType,
	)
	if err != nil {
		os.Remove(filePath) // Clean up if database operation fails
		return nil, fmt.Errorf("failed to save image metadata: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	// Get the created record
	var image ImageMetadata
	err = db.QueryRow(
		`SELECT id, user_id, original_filename, stored_filename, file_path, file_size, mime_type, created_at 
		FROM user_images WHERE id = ?`, id,
	).Scan(
		&image.ID, &image.UserID, &image.OriginalFilename, &image.StoredFilename,
		&image.FilePath, &image.FileSize, &image.MIMEType, &image.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch saved image metadata: %v", err)
	}

	return &image, nil
}

// GetUserImages returns all images uploaded by a specific user
func GetUserImages(userID int) ([]ImageMetadata, error) {
	rows, err := db.Query(
		`SELECT id, user_id, original_filename, stored_filename, file_path, file_size, mime_type, created_at 
		FROM user_images WHERE user_id = ? ORDER BY created_at DESC`, userID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query user images: %v", err)
	}
	defer rows.Close()

	var images []ImageMetadata
	for rows.Next() {
		var img ImageMetadata
		err := rows.Scan(
			&img.ID, &img.UserID, &img.OriginalFilename, &img.StoredFilename,
			&img.FilePath, &img.FileSize, &img.MIMEType, &img.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan image row: %v", err)
		}
		images = append(images, img)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over image rows: %v", err)
	}

	return images, nil
}

// DeleteImage deletes an image and its metadata
func DeleteImage(userID, imageID int) error {
	// Get the file path first
	var filePath string
	err := db.QueryRow(
		"SELECT file_path FROM user_images WHERE id = ? AND user_id = ?",
		imageID, userID,
	).Scan(&filePath)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("image not found or access denied")
		}
		return fmt.Errorf("failed to get image path: %v", err)
	}

	// Delete from database
	result, err := db.Exec(
		"DELETE FROM user_images WHERE id = ? AND user_id = ?",
		imageID, userID,
	)
	if err != nil {
		return fmt.Errorf("failed to delete image metadata: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("image not found or access denied")
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		// Log the error but don't fail the operation if file deletion fails
		log.Printf("Warning: failed to delete image file %s: %v", filePath, err)
	}

	return nil
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
