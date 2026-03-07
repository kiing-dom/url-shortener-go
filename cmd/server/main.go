package main

import (
	"fmt"
	"net/http"

	"github.com/kiing-dom/url-shortener-go/internal/handler"
	"github.com/kiing-dom/url-shortener-go/internal/repository"
	"github.com/kiing-dom/url-shortener-go/internal/service"
)

func main() {
	fmt.Println("starting server...")

	repo := repository.NewInMemoryURLRepository()
	svc := service.NewURLService(repo)
	handler := handler.NewURLHandler(svc)

	// wiring handler methods
	http.HandleFunc("/shorten", handler.HandleShorten)
	http.HandleFunc("/", handler.HandleRedirect)
	http.HandleFunc("/stats/", handler.HandleStats)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// setting up port to listen and serve
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed", err)
	}
}
