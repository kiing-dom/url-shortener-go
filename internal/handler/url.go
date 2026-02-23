package handler

import (
	"github.com/kiing-dom/url-shortener-go/internal/service"
	"net/http"
)

type URLHandler struct {
	// TODO: create the URLService
	service *service.URLService
}

func NewURLHandler(svc *service.URLService) *URLHandler {
	return &URLHandler{
		service: svc,
	}
}

func (h *URLHandler) HandleShorten(w http.ResponseWriter, r *http.Request) {
	// TODO: complete the function
}

func (h *URLHandler) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// TODO: complete the function
}
