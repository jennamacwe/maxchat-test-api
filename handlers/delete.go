package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"pdf-api/config"
	"pdf-api/utils"
)

func DeletePDF(w http.ResponseWriter, r *http.Request) {
	// mengambil ID dari URL
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	// expected: ["api", "pdf", "{id}"]
	if len(parts) != 3 {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	// status file
	var status string
	err = config.DB.QueryRow(
		"SELECT status FROM pdf_files WHERE id = ?",
		id,
	).Scan(&status)

	if err != nil {
		utils.JSON(w, http.StatusNotFound, utils.APIResponse{
			Success: false,
			Message: "File not found",
		})
		return
	}

	if status == "DELETED" {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "File already deleted",
		})
		return
	}

	res, err := config.DB.Exec(
		"UPDATE pdf_files SET status = 'DELETED', deleted_at = ? WHERE id = ?",
		time.Now(),
		id,
	)

	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to delete file",
		})
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: "Failed to check affected rows",
		})
		return
	}

	if rows == 0 {
		utils.JSON(w, http.StatusNotFound, utils.APIResponse{
			Success: false,
			Message: "No data updated",
		})
		return
	}


	utils.JSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "PDF deleted successfully",
		Data: map[string]interface{}{
			"id":         id,
			"status":     "DELETED",
			"deleted_at": time.Now(),
		},
	})
}