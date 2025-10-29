package main

import (
	"fmt"
	"log"
	"net/http"
	"web-analyzer/internal/delivery/handler"
	"web-analyzer/internal/service"
	"web-analyzer/internal/analyzer"
)

func main() {
	analyzerService := service.NewAnalyzerService(analyzer.NewHTMLAnalyzer())
	handler := handler.NewHandler(analyzerService)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler.Router()))
}
