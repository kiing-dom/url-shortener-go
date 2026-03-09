package repository

import (
	"github.com/kiing-dom/url-shortener-go/internal/model"
)

type URLRepository interface {
	Save(url, code string) error
	FindByURL(url string) (string, error)
	FindByCode(code string) (string, error)
	FindEntryByCode(code string) (*model.URLEntry, error)
	IncrementClicks(code string) error
}
