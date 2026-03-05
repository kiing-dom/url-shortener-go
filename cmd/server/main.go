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

	// setting up port to listen and serve
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed", err)
	}
}
