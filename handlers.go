package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// getMovies handles GET /movies
func getMovies(c *gin.Context) {
	var movies []Movie

	// Find runs: SELECT * FROM movies
	err := db.Find(&movies).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, movies)
}

// getMovieByID handles GET /movies/:id
func getMovieByID(c *gin.Context) {
	// GORM treats a number as a primary key, but a raw string would be treated as SQL.
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var movie Movie

	err = db.First(&movie, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, movie)
}

// createMovie handles POST /movies
func createMovie(c *gin.Context) {
	var movie Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validateMovie(movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Create(&movie).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, movie)
}

// validateMovie checks a movie's fields before it is saved.
func validateMovie(m Movie) error {
	if strings.TrimSpace(m.Title) == "" {
		return errors.New("title is required")
	}

	if m.Rating != nil && (*m.Rating < 0 || *m.Rating > 10) {
		return errors.New("rating must be between 0 and 10")
	}

	// 1888 is the year of the first known film.
	if m.ReleaseYear != nil && *m.ReleaseYear < 1888 {
		return errors.New("release_year must be 1888 or later")
	}

	return nil
}

// home handles GET /
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Movies API",
		"endpoints": []string{
			"GET  /health",
			"GET  /movies",
			"GET  /movies/:id",
			"POST /movies",
		},
	})
}

// It pings Postgres, so "ok" means both the server and the database are up.
func health(c *gin.Context) {
	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "database": err.Error()})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "database": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "connected"})
}
