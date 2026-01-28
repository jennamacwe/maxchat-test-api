# PDF REST API (Golang)

Backend REST API untuk **upload, generate, list, dan delete PDF** menggunakan **Golang (tanpa framework)** dan **MySQL**. Project ini dibuat sebagai bagian dari **Technical Test Backend Developer – PT Maxchat Indonesia**.

---

## Tech Stack

- **Golang** v1.25.6
- **MySQL**
- **gofpdf** (PDF Generator)
- **Postman** (API Testing)
- **Laragon** (Local Server – Windows)

---

## Project Structure

```
pdf-api/
│
├── config/               # Database configuration
│   └── database.go
│
├── handlers/             # HTTP handlers
│   ├── upload.go
│   ├── generate.go
│   ├── list.go
│   └── delete.go
│
├── utils/                # Helper & utilities
│   ├── response.go       # JSON response helper
│   └── pdf.go            # PDF generation logic
│
├── uploads/
│   └── pdf/              # Uploaded & generated PDF files
│
├── postman/
│   └── pdf-api.postman_collection.json
│
├── main.go               # Application entry point
├── go.mod
└── README.md
```

---

## Setup & Installation

### 1. Clone Repository

```bash
git clone https://github.com/jennamacwe/maxchat-test-api.git

cd pdf-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Database

Buat database MySQL:

```sql
CREATE DATABASE pdf_api;

USE pdf_api;

CREATE TABLE pdf_files (
  id INT AUTO_INCREMENT PRIMARY KEY,
  filename VARCHAR(255),
  filepath TEXT,
  status VARCHAR(50),
  created_at DATETIME
);
```

Atur konfigurasi database di:

```
config/database.go
```

---

## Run Application

```bash
go run main.go
```

Server berjalan di:
```
http://localhost:8080
```

Health Check:
```
GET /api/status
```

---

## API Endpoints

### 1. Upload PDF

```
POST /api/pdf/upload
```

**Body (form-data):**
| Key  | Type | Required |
|------|------|----------|
| file | File |    ✅    |

---

### 2. Generate PDF

```
POST /api/pdf/generate
```

**Headers:**
```
Content-Type: application/json
```

**Body (JSON):**
```json
{
  "title": "PDF Dengan Logo",
  "institution_name": "PT Maxchat Indonesia",
  "address": "Jakarta",
  "phone": "021-123456",
  "logo_url": "default",
  "content": "Ini PDF sudah ada logo di header kiri."
}
```

---

### 3. List PDF

```
GET /api/pdf/list
```

---

### 4. Delete PDF

```
DELETE /api/pdf/{id}
```

---

## Testing

Semua endpoint diuji menggunakan **Postman**:
- JSON request
- Multipart form-data
- Error handling

---

## Features Implemented

- REST API tanpa framework
- Upload PDF
- Generate PDF dari JSON
- Header & Footer PDF
- Simpan file ke local storage
- Simpan metadata ke MySQL
- Soft delete data
- Error handling & validation

---

## Author

**Jennatul Macwe**  
Backend Developer Candidate  

---

## Notes

Project ini dibuat sesuai dengan kebutuhan technical test dan dapat dikembangkan lebih lanjut dengan:
- Authentication
- Pagination
- Cloud storage
- Logging middleware

---

Terima kasih atas kesempatan yang diberikan

Best regards,
Jennatul Macwe

