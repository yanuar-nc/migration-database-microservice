package domain

import "time"

type Migration struct {
	ID        int       `json:"id"`
	Version   int64     `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
