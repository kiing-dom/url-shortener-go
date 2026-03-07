package service

import (
	"errors"
	"math/rand/v2"
	"net/url"

	"github.com/kiing-dom/url-shortener-go/internal/model"
	"github.com/kiing-dom/url-shortener-go/internal/repository"
)

type URLService struct {
	repo repository.URLRepository
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewURLService(r repository.URLRepository) *URLService {
	return &URLService{
		repo: r,
	}
}

func (s *URLService) Shorten(input string) (string, error) {
	// validate input URL
	if err := validateURL(input); err != nil {
		return "", err
	}

	if existingCode, err := s.repo.FindByURL(input); err == nil {
		return existingCode, nil
	}

	// generate unique code
	code := generateCode()

	// save mapping of code-URL to repo. if it exists already return error
	if err := s.repo.Save(input, code); err != nil {
		return "", err
	}

	return code, nil
}

func (s *URLService) Resolve(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return "", errors.New("code not found")
	}

	s.repo.IncrementClicks(code)

	return url, nil
}

func (s *URLService) GetStats(code string) (*model.URLEntry, error) {
	return s.repo.FindEntryByCode(code)
}

func generateCode() string {
	// generate an 7-char code that can use a-zA-Z0-9
	code := make([]byte, 7)
	for i := range code {
		code[i] = charset[rand.N(len(charset))]
	}

	return string(code)
}

func validateURL(input string) error {
	// validatiton logic
	parsedURL, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("invalid URL structure")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL must have http or https scheme")
	}

	return nil
}
