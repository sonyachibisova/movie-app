package models

type Movie struct {
	ID          uint    `gorm:"promarykey" json:"id"`
	Title       string  `gorm:"not null" json:"title"`
	Description string  `json:"description"`
	Director    string  `json:"director"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	PosterURL   string  `json:"poster_url"`
	AvgRating   float32 `gorm:"default:0" json:"avg_rating"`
	Country     string  `json:"country"`
	Duration    int     `json:"duration"`
	Actors      string  `json:"actors"`
}
