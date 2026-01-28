package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
	"time"
	"os"
	"fmt"

	"pdf-api/config"
	"pdf-api/utils"
)

func UploadPDF(w http.ResponseWriter, r *http.Request) {
	// membatasi ukuran max 10MB
	const maxFileSize = 10 << 20 // 10MB
	file, handler, err := r.FormFile("file")

	if handler.Size > maxFileSize {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "File size exceeds maximum limit (10MB)",
		})
		return
	}

	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "File is required",
		})
		return
	}
	defer file.Close()

	// Validasi ekstensi
	ext := strings.ToLower(filepath.Ext(handler.Filename))
	if ext != ".pdf" {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Only PDF files are allowed",
		})
		return
	}

	// Validasi MIME type
	if handler.Header.Get("Content-Type") != "application/pdf" {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Invalid MIME type",
		})
		return
	}

	// Nama file unik
	filename := fmt.Sprintf(
		"upload_%d%s",
		time.Now().UnixNano(),
		ext,
	)

	uploadPath := "uploads/pdf/" + filename

	// Pastikan folder ada
	os.MkdirAll("uploads/pdf", os.ModePerm)

	dst, err := os.Create(uploadPath)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}
	defer dst.Close()

	_, err = dst.ReadFrom(file)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to write file",
		})
		return
	}

	// Simpan ke database
	query := `
		INSERT INTO pdf_files
		(filename, original_name, filepath, size, status, created_at)
		VALUES (?, ?, ?, ?, 'UPLOADED', ?)
	`

	result, err := config.DB.Exec(
		query,
		filename,
		handler.Filename,
		uploadPath,
		handler.Size,
		time.Now(),
	)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to save data",
		})
		return
	}

	id, _ := result.LastInsertId()

	utils.JSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "PDF uploaded successfully",
		Data: map[string]interface{}{
			"id":            id,
			"original_name": handler.Filename,
			"filename":      filename,
			"filepath":      uploadPath,
			"size":          handler.Size,
			"status":        "UPLOADED",
			"created_at":    time.Now(),
		},
	})
}