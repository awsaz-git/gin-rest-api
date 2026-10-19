package main

import "time"

// The struct Movie maps to the table "movies",
type Movie struct {
	ID          int64     `json:"id"           gorm:"primaryKey"`
	Title       string    `json:"title"        gorm:"not null"`
	Genre       *string   `json:"genre"`
	ReleaseYear *int      `json:"release_year"`
	Rating      *float64  `json:"rating"       gorm:"type:numeric(3,1)"`
	Watched     bool      `json:"watched"      gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at"`
}
