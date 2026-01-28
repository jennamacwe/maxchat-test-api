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
cd maxchat-test-api
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
```

Create Table:
```sql
CREATE TABLE pdf_files (
  id BIGINT NOT NULL AUTO_INCREMENT,
  filename VARCHAR(255) NOT NULL,
  original_name VARCHAR(255) DEFAULT NULL,
  filepath VARCHAR(500) NOT NULL,
  size BIGINT DEFAULT NULL,
  status ENUM('CREATED','UPLOADED','DELETED') NOT NULL DEFAULT 'CREATED',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (id)
);
```

Atur konfigurasi database di:

```
config/database.go
```
Sesuaikan:
- host
- username
- password
- database name

---

## Run Application

```bash
go run main.go
```

Server berjalan di:
```
http://localhost:8080
```

Status Check:
```
GET /api/status
```

---

## API Endpoints
Base URL
```
http://localhost:8080/api
```

### 1. Upload PDF
**Endpoint**:

```
POST /api/pdf/upload
```

**Headers**:
```
Content-Type: multipart/form-data
```

**Body (form-data):**
| Key  | Type | Required |
|------|------|----------|
| file | File |    ✅    |

**Success Response**:
 ```json
{
  "success": true,
  "message": "PDF uploaded successfully",
  "data": {
    "id": 1,
    "original_name": "document.pdf",
    "filename": "upload_1700000000000.pdf",
    "filepath": "uploads/pdf/upload_1700000000000.pdf",
    "size": 102400,
    "status": "UPLOADED",
    "created_at": "2026-01-28T10:00:00Z"
  }
}
```

**Error Response**:
1. File tidak dikirim

```json
{
  "success": false,
  "message": "File is required"
}
```
Terjadi jika field file tidak ada di form-data.

2. Ukuran file melebihi 10MB

```json
{
  "success": false,
  "message": "File size exceeds maximum limit (10MB)"
}
```
Maksimal ukuran file adalah 10MB.

3. File bukan PDF (ekstensi salah)

```json
{
  "success": false,
  "message": "Only PDF files are allowed"
}
```
Ekstensi selain .pdf akan ditolak.

4. MIME type tidak valid

```json
{
  "success": false,
  "message": "Invalid MIME type"
}
```

Content-Type file harus application/pdf.

5. Gagal menyimpan file ke server

```json
{
  "success": false,
  "message": "Failed to save file"
}
```

Masalah permission atau folder tidak bisa dibuat.

6. Gagal menulis file

```json
{
  "success": false,
  "message": "Failed to write file"
}
```

Error saat proses penulisan file ke storage.

7. Gagal menyimpan data ke database
   
```json
{
  "success": false,
  "message": "Failed to save data"
}
```

Error saat insert data ke tabel pdf_files.

---

### 2. Generate PDF
**Endpoint**:

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

**Success Response**:
```json
{
  "success": true,
  "message": "PDF generated successfully",
  "data": {
    "filename": "generated_1700000000.pdf"
  }
}
```
---

### 3. List PDF
**Endpoint**:

```
GET /api/pdf/list
```

**Success Response**:
```json
{
  "success": true,
  "message": "PDF list retrieved",
  "data": [
    {
      "id": 1,
      "filename": "file_1700000000.pdf",
      "status": "UPLOADED",
      "created_at": "2026-01-28 10:00:00"
    }
  ]
}
```

---

### 4. Delete PDF (Soft Delete)
**Endpoint**:

```
DELETE /api/pdf/{id}
```

**Success Response**:
```json
{
  "success": true,
  "message": "PDF deleted successfully",
  "data": {
    "id": 1,
    "status": "DELETED",
    "deleted_at": "2026-01-28 10:10:00"
  }
}
```

**Error Response**:
```json
{
  "success": false,
  "message": "File not found"
}
```
---

## Database Schema Documentation
| Column Name     | Type      | Description                      |
| --------------- | --------- | -------------------------------- |
| `id`            | BIGINT    | Primary Key (Auto Increment)     |
| `filename`      | VARCHAR   | Stored filename                  |
| `original_name` | VARCHAR   | Original uploaded filename       |
| `filepath`      | VARCHAR   | File path on server              |
| `size`          | BIGINT    | File size (bytes)                |
| `status`        | ENUM      | `CREATED`, `UPLOADED`, `DELETED` |
| `created_at`    | TIMESTAMP | Created time                     |
| `updated_at`    | TIMESTAMP | Last updated time                |
| `deleted_at`    | TIMESTAMP | Soft delete timestamp            |

---

## Testing

Semua endpoint diuji menggunakan **Postman**:
- JSON request
- Multipart form-data
- Error handling

**Postman Collection tersedia di folder:**
```json
postman/pdf-api.postman_collection.json
```

Berisi:
- Semua endpoint
- Contoh request & response
- Environment variable

---

## Features Implemented

- REST API tanpa framework
- Upload & generate PDF
- Metadata storage ke MySQL
- Soft delete implementation
- Validasi & error handling
- Clean project structure
- Postman documentation

---

## Author

**Jennatul Macwe**  
Backend Developer Candidate  

---

Terima kasih atas kesempatan yang diberikan

Best regards,
Jennatul Macwe
