package domain

import "time"

type Movie struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Rating      float32   `json:"rating"`
	Awards      int       `json:"awards"`
	ReleaseDate string    `json:"release_date"`
	Length      *int      `json:"length"`
	GenreID     *int      `json:"genre_id"`
}
