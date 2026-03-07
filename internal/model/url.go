package model

import "time"

type URLEntry struct {
	Code        string
	OriginalURL string
	CreatedAt   time.Time
	Clicks      int
}
