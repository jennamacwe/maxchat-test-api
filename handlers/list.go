package handlers

import (
	"net/http"
	"strconv"

	"pdf-api/config"
	"pdf-api/models"
	"pdf-api/utils"
)

func ListPDF(w http.ResponseWriter, r *http.Request) {
	// Ambil query param
	status := r.URL.Query().Get("status")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Query dasar
	query := `
		SELECT id, filename, original_name, size, status, created_at
		FROM pdf_files
		WHERE 1=1
	`

	args := []interface{}{}

	// Filter status
	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}

	// Urut & pagination
	query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to fetch data",
		})
		return
	}
	defer rows.Close()

	files := []models.PDFFile{}

	for rows.Next() {
		var file models.PDFFile
		err := rows.Scan(
			&file.ID,
			&file.Filename,
			&file.OriginalName,
			&file.Size,
			&file.Status,
			&file.CreatedAt,
		)
		if err != nil {
			continue
		}
		files = append(files, file)
	}

	// Hitung total data
	var total int
	countQuery := "SELECT COUNT(*) FROM pdf_files"
	if status != "" {
		countQuery += " WHERE status = ?"
		err = config.DB.QueryRow(countQuery, status).Scan(&total)
	} else {
		err = config.DB.QueryRow(countQuery).Scan(&total)
	}

	utils.JSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"data": files,
			"pagination": map[string]interface{}{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	})
}
