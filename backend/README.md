# Dolil Lekhok Backend

This is the backend service for Dolil Lekhok, built with Go and Gin framework.

## Setup

1. Make sure you have Go installed (version 1.16 or higher)
2. Clone the repository
3. Navigate to the backend directory
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Run the server:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Check if the server is running

### Authentication
- `POST /api/auth/signup` - Register a new user
  ```json
  {
    "username": "testuser",
    "password": "password123"
  }
  ```

- `POST /api/auth/login` - Login with existing user
  ```json
  {
    "username": "testuser",
    "password": "password123"
  }
  ```

## Development

- The server automatically reloads with [air](https://github.com/cosmtrek/air) for development
- To install air:
  ```bash
  go install github.com/cosmtrek/air@latest
  ```
- Then run:
  ```bash
  air
  ```
