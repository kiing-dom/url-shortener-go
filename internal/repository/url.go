package repository

import (
	"errors"
	"sync"
)

type URLRepository interface {
	Save(url, code string) error
	FindByURL(url string) (string, error)
	FindByCode(code string) (string, error)
	IncrementClicks(code string) error
}

type InMemoryURLRepository struct {
	mu        sync.RWMutex
	codeToURL map[string]string // map[code]url
	URLToCode map[string]string // map[url]code
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		codeToURL: make(map[string]string),
		URLToCode: make(map[string]string),
	}
}

// Save stores the URL and code in memory
func (r *InMemoryURLRepository) Save(url, code string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.codeToURL[code]; exists {
		return errors.New("code already exists")
	}

	if _, exists := r.URLToCode[url]; exists {
		return errors.New("URL already exists")
	}

	r.codeToURL[code] = url
	r.URLToCode[url] = code
	return nil
}

// FindByURL retrieves code for a given URL
func (r *InMemoryURLRepository) FindByURL(url string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for code, storedURL := range r.codeToURL {
		if storedURL == url {
			return code, nil
		}
	}

	return "", errors.New("URL not found")
}

// FindByCode retrieves URL for a given code
func (r *InMemoryURLRepository) FindByCode(code string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, exists := r.codeToURL[code]
	if !exists {
		return "", errors.New("code not found")
	}

	return url, nil
}

// placeholder for incrementing clicks
func (r *InMemoryURLRepository) IncrementClicks(code string) error {
	// TODO: finish impl
	return nil
}
