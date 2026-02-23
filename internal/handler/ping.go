package handler

import (
	"fmt"
	"net/http"
)

type PingHandler struct {
	AppName string
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[%s] pong!", h.AppName)
	fmt.Println("sent ping to server!")
}
