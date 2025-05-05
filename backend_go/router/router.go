package router

import (
	"backend_go/db"
	"backend_go/db/dao"
	"backend_go/router/rest"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	// ✅ Explicitly define CORS rules
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache pre-flight OPTIONS request
	}))

	// ✅ Handle OPTIONS preflight requests
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204) // No Content
	})

	// ✅ Debugging middleware (optional, remove in production)
	r.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
	})

	// Connect to the database
	database, err := db.GetDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Initialize DAOs
	userDAO := dao.NewUserDAO(database)
	expenseDAO := dao.NewHaushaltsausgabenDAO(database)

	// General route to test if server is running
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Register routes
	rest.RegisterUserRoutes(r, userDAO)
	rest.RegisterHaushaltsausgabenRoutes(r, expenseDAO)

	return r
}
