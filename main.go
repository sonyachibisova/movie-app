package main

import (
	"movie-app/database"
	"movie-app/models"

	"github.com/gin-gonic/gin"
)

func stringToUint(s string) uint {
	var i uint
	for _, c := range s {
		i = i*10 + uint(c-'0')
	}
	return i
}

func stringToFloat(s string) float32 {
	var f float32
	for i, ch := range s {
		if ch == '.' {
			continue
		}

		f = f*10 + float32(ch-'0')
		if i == 0 && len(s) > 1 && s[1] == '.' {
			f = f / 10
		}
	}

	return f
}

func updateAvgRating(movieID uint) {
	var ratings []float32
	database.DB.Model(&models.UserMovie{}).
		Where("movie_id = ? AND status = 'watched' AND rating is NOT NULL", movieID).
		Pluck("rating", &ratings)

	var sum float32
	for _, r := range ratings {
		sum += r
	}

	avg := float32(0)
	if len(ratings) > 0 {
		avg = sum / float32(len(ratings))
	}

	database.DB.Model(&models.Movie{}).
		Where("id = ?", movieID).Update("avg_rating", avg)
}

func main() {
	database.InitDB()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		var movies []models.Movie
		database.DB.Find(&movies)
		c.HTML(200, "index.html", gin.H{"movies": movies})
	})

	r.GET("/profile", func(c *gin.Context) {
		var user models.User
		database.DB.Where("username = ?", "me").First(&user)

		var UserMovies []models.UserMovie
		database.DB.Preload("Movie").Where("user_id = ?", user.ID).Find(&UserMovies)

		var watched, want []models.UserMovie
		for _, um := range UserMovies {
			if um.Status == "watched" {
				watched = append(watched, um)
			} else {
				want = append(want, um)
			}
		}

		c.HTML(200, "profile.html", gin.H{
			"watched": watched,
			"want":    want,
		})
	})

	r.POST("/movies/:id/:status", func(c *gin.Context) {
		movieID := c.Param("id")
		status := c.Param("status")

		var user models.User
		database.DB.Where("username = ?", "me").First(&user)

		userMovie := models.UserMovie{
			UserID:  user.ID,
			MovieID: stringToUint(movieID),
			Status:  status,
		}

		result := database.DB.Where("user_id = ? AND movie_id = ?",
			user.ID, movieID).FirstOrCreate(&userMovie)

		if result.Error != nil {
			c.String(500, "Ошибка")
			return
		}

		c.String(200, "OK")
	})

	r.POST("/movies/:id/rate", func(c *gin.Context) {
		movieID := c.Param("id")
		rating := c.PostForm("rating")

		var user models.User
		database.DB.Where("username = ?", "me").First(&user)

		var userMovie models.UserMovie
		result := database.DB.Where("user_id = ? AND movie_id = ? AND status = ?", user.ID, movieID).
			First(&userMovie)

		if result.Error != nil {
			c.String(404, "Фильм не найден в просмотренных")
			return
		}

		var ratingFloat float32
		ratingFloat = stringToFloat(rating)
		userMovie.Rating = &ratingFloat
		database.DB.Save(&userMovie)

		updateAvgRating(stringToUint(movieID))

		c.String(200, "OK")
	})

	r.Run(":8081")
}
