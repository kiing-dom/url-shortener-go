package repository

type URLRepository interface {
	Save(url, code string) error
	FindByURL(url string) (string, error)
	FindByCode(code string) (string, error)
	IncrementClicks(code string) error
}
