# Dolil Lekhok

**Dolil Lekhok** is a micro SaaS platform that converts Bangla handwritten documents into structured digital formats (PDF, Word, etc.). This project uses **SvelteKit** for the frontend and **Go (Gin)** for the backend API.

---

## **Table of Contents**

- [Dolil Lekhok](#dolil-lekhok)
  - [**Table of Contents**](#table-of-contents)
  - [**Features**](#features)
  - [**Tech Stack**](#tech-stack)
  - [**Prerequisites**](#prerequisites)
  - [**Getting Started**](#getting-started)
    - [**1. Clone the Repository**](#1-clone-the-repository)
    - [**2. Setup Backend (Go + Gin)**](#2-setup-backend-go--gin)
    - [**3. Setup Frontend (SvelteKit with pnpm)**](#3-setup-frontend-sveltekit-with-pnpm)
  - [**Project Structure**](#project-structure)
  - [**Running the Project Together**](#running-the-project-together)
  - [**Contributing**](#contributing)

---

## **Features**

* Upload handwritten Bangla images.
* Convert handwriting to digital text using OCR.
* Generate formatted PDFs or Word documents.
* Async processing for large files.
* Lightweight and deployable on budget VPS.

---

## **Tech Stack**

* **Frontend:** SvelteKit, TypeScript, pnpm
* **Backend:** Go + Gin framework
* **OCR Engine:** Tesseract or Python microservice (optional for advanced handwriting recognition)
* **PDF Generation:** Go libraries (`pdfcpu`, `unidoc`)
* **Job Queue:** Redis / Go worker pool (optional)
* **Database:** SQLite / PostgreSQL (optional for persistence)

---

## **Prerequisites**

* **Node.js:** v18+
* **pnpm** package manager
* **Go:** v1.20+
* **Tesseract OCR:** Latest version installed and in PATH (for CLI integration)
* Optional: Python 3.x + dependencies if using Python microservice for OCR
* Redis (optional for async queue)

---

## **Getting Started**

### **1. Clone the Repository**

```bash
git clone https://github.com/yourusername/dolil-lekhok.git
cd dolil-lekhok
```

---

### **2. Setup Backend (Go + Gin)**

1. Navigate to backend folder:

```bash
cd backend
```

2. Install dependencies:

```bash
go mod tidy
```

3. Start backend server:

```bash
go run main.go
```

* API will run on `http://localhost:8080`.
* **Note:** Implement asynchronous OCR processing using a worker queue to avoid blocking Gin requests.

---

### **3. Setup Frontend (SvelteKit with pnpm)**

1. Navigate to frontend folder:

```bash
cd ../frontend
```

2. Install dependencies:

```bash
pnpm install
```

3. Run the development server:

```bash
pnpm dev -- --open
```

* Frontend will run on `http://localhost:5173` (or the URL shown in console).

---

## **Project Structure**

```
dolil_lekhok/
├── backend/             # Go backend (Gin framework)
│   ├── main.go          # Entry point
│   ├── routes/          # API route handlers
│   ├── controllers/     # Business logic
│   ├── models/          # Data models
│   ├── services/        # OCR & PDF generation
│   └── utils/           # Helper functions
├── frontend/            # SvelteKit frontend
│   ├── src/
│   │   ├── routes/      # SvelteKit pages
│   │   ├── lib/         # Components & utilities
│   │   └── assets/      # Static assets
├── README.md
└── .gitignore
```

---

## **Running the Project Together**

1. Start backend server (`localhost:8080`).
2. Start frontend server (`localhost:5173`).
3. Open the frontend in your browser.
4. Upload handwritten images and test PDF generation.
5. For large uploads or slow OCR, results can be handled asynchronously with a queue and job IDs.

## **Contributing**

* Fork the repo.
* Create a feature branch: `git checkout -b feature/your-feature`
* Commit changes: `git commit -m "Add feature"`
* Push to branch: `git push origin feature/your-feature`
* Open a Pull Request

---

✅ **Notes / Tips for Your VPS Setup**

* Use a **worker queue** (Go channels or Redis) to process OCR asynchronously, preventing Gin from blocking.
* Preprocess images (resize, grayscale) before OCR to reduce CPU load.
* Cache repeated OCR results for duplicate images.
* Batch multiple images per PDF to improve throughput.