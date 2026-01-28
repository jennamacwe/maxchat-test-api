package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"fmt"

	"pdf-api/config"
	"pdf-api/utils"
)


type GeneratePDFRequest struct {
	Title           string `json:"title"`
	InstitutionName string `json:"institution_name"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	LogoURL         string `json:"logo_url"`
	Content         string `json:"content"`
}

func GeneratePDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSON(w, http.StatusMethodNotAllowed, utils.APIResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var req GeneratePDFRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Invalid JSON body",
		})
		return
	}

	if req.Title == "" || req.InstitutionName == "" || req.Content == "" {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Required fields are missing",
		})
		return
	}

	filename := fmt.Sprintf("report_%d.pdf", time.Now().UnixNano())
	filePath := "uploads/pdf/" + filename
	os.MkdirAll("uploads/pdf", os.ModePerm)

	logoPath := "assets/logo/default.png"

	if req.LogoURL != "" {
		logoPath = "assets/logo/logo.png"
	}

	pdfData := utils.PDFContent{
		Title:           req.Title,
		InstitutionName: req.InstitutionName,
		Address:         req.Address,
		Phone:           req.Phone,
		Content:         req.Content,
		LogoPath:        logoPath,
	}

	err = utils.GenerateFullPDF(filePath, pdfData)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to generate PDF",
		})
		return
	}

	// SAVE TO DATABASE
	query := `
		INSERT INTO pdf_files
		(filename, filepath, status, created_at)
		VALUES (?, ?, 'CREATED', ?)
	`

	result, err := config.DB.Exec(query, filename, filePath, time.Now())
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
		Message: "PDF generated successfully",
		Data: map[string]interface{}{
			"id":         id,
			"filename":   filename,
			"filepath":   filePath,
			"status":     "CREATED",
			"created_at": time.Now(),
		},
	})
}
