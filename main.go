package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Load DATABASE_URL from .env into the environment.
	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")

	var err error
	db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to Postgres: ", err)
	}

	// AutoMigrate creates the movies table from the Movie struct if it doesn't exist,
	db.AutoMigrate(&Movie{})

	router := gin.Default()
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovieByID)
	router.POST("/movies", createMovie)

	router.Run("localhost:8080")
}
