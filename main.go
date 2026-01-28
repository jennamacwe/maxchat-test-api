package main

import (
	"log"
	"net/http"

	"pdf-api/config"
	"pdf-api/handlers"
	"pdf-api/utils"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/api/status", apiStatus)
	http.HandleFunc("/api/pdf/upload", handlers.UploadPDF)
	http.HandleFunc("/api/pdf/list", handlers.ListPDF)
	http.HandleFunc("/api/pdf/generate", handlers.GeneratePDF)
	http.HandleFunc("/api/pdf/", handlers.DeletePDF)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiStatus(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "PDF API is running",
	})
}