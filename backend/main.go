package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zven/bookpavilion/config"
	"github.com/zven/bookpavilion/controllers"
	"github.com/zven/bookpavilion/services"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize services
	bookService := services.NewBookService(config.GetDB())

	// Initialize controllers
	bookController := controllers.NewBookController(bookService)

	// Set up Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Set maximum multipart form size (default is 32 MB)
	r.MaxMultipartMemory = 8 << 20 // 8 MB

	// Serve static files (uploaded books)
	r.Static("/uploads", config.GetUploadDir())

	// API routes
	api := r.Group("/api")
	{
		// Book routes
		books := api.Group("/books")
		{
			books.POST("", bookController.CreateBook)
			books.GET("", bookController.ListBooks)
			books.GET("/:id", bookController.GetBook)
			books.DELETE("/:id", bookController.DeleteBook)
			books.GET("/:id/content", bookController.GetBookContent) // Add this line
		}

		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
