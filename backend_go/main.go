package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default()) // Add CORS middleware
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from Go REST"})
	})
	r.Run(":8080")
}
