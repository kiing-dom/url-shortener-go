package main

import (
	"fmt"
	"net/http"

	"github.com/kiing-dom/url-shortener-go/internal/handler"
)

func main() {
	fmt.Println("starting server...")

	ph := &handler.PingHandler{
		AppName: "Go-URL-Shortener-V1",
	}

	http.Handle("/ping", ph)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed!", err)
	}
}
