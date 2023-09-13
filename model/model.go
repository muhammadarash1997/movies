package model

import (
	"time"
)

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s Movie) Validate() bool {
	return !(s.IsEmpty() || s.ID == 0 || s.Title == "" || s.Description == "" || s.Rating < 0 || s.Image == "")
}

func (s Movie) IsEmpty() bool {
	return s == Movie{}
}
