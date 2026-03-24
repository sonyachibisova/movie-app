package models

type Movie struct {
	ID          uint   `gorm:"promarykey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	PosterURL   string `json:"poster_url"`
}
