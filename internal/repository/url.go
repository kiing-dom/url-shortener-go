package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/kiing-dom/url-shortener-go/internal/model"
)

type URLRepository interface {
	Save(url, code string) error
	FindByURL(url string) (string, error)
	FindByCode(code string) (string, error)
	FindEntryByCode(code string) (*model.URLEntry, error)
	IncrementClicks(code string) error
}

type InMemoryURLRepository struct {
	mu          sync.RWMutex
	codeToEntry map[string]*model.URLEntry // map[code]URLEntry
	urlToCode   map[string]string          // map[url]code
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		codeToEntry: make(map[string]*model.URLEntry),
		urlToCode:   make(map[string]string),
	}
}

const noCode = "code not found"

// Save stores the URL and code in memory
func (r *InMemoryURLRepository) Save(url, code string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.codeToEntry[code]; exists {
		return errors.New("code already exists")
	}

	if _, exists := r.urlToCode[url]; exists {
		return errors.New("URL already exists")
	}

	entry := &model.URLEntry{
		Code:        code,
		OriginalURL: url,
		CreatedAt:   time.Now(),
		Clicks:      0,
	}

	r.codeToEntry[code] = entry
	r.urlToCode[url] = code
	return nil
}

// FindByURL retrieves code for a given URL
func (r *InMemoryURLRepository) FindByURL(url string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	code, exists := r.urlToCode[url]
	if !exists {
		return "", errors.New("URL not found")
	}

	return code, nil
}

// FindByCode retrieves URL for a given code
func (r *InMemoryURLRepository) FindByCode(code string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entry, exists := r.codeToEntry[code]
	if !exists {
		return "", errors.New(noCode)
	}

	return entry.OriginalURL, nil
}

func (r *InMemoryURLRepository) FindEntryByCode(code string) (*model.URLEntry, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entry, exists := r.codeToEntry[code]
	if !exists {
		return nil, errors.New(noCode)
	}

	return entry, nil
}

// placeholder for incrementing clicks
func (r *InMemoryURLRepository) IncrementClicks(code string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	entry, exists := r.codeToEntry[code]
	if !exists {
		return errors.New(noCode)
	}

	entry.Clicks++
	return nil
}
