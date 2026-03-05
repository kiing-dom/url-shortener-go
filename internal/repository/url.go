package repository

import "errors"

type URLRepository interface {
	Save(url, code string) error
	FindByURL(url string) (string, error)
	FindByCode(code string) (string, error)
	IncrementClicks(code string) error
}

type InMemoryURLRepository struct {
	data map[string]string // map[code]url
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		data: make(map[string]string),
	}
}

// Save stores the URL and code in memory
func (r *InMemoryURLRepository) Save(url, code string) error {
	if _, exists := r.data[code]; exists {
		return errors.New("code already exists")
	}

	r.data[code] = url
	return nil
}

// FindByURL retrieves code for a given URL
func (r *InMemoryURLRepository) FindByURL(url string) (string, error) {
	for code, storedURL := range r.data {
		if storedURL == url {
			return code, nil
		}
	}

	return "", errors.New("URL not found")
}

// FindByCode retrieves URL for a given code
func (r *InMemoryURLRepository) FindByCode(code string) (string, error) {
	url, exists := r.data[code]
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
