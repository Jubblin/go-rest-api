package main

import (
	"fmt"
	"go-rest-api/controllers"
	_ "go-rest-api/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Book API
// @version         1.0
// @description     A simple book management API.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {

	for _ , arg := range os.Args {
		if arg == "healthcheck" {
			fmt.Println("OK")
			os.Exit(0)
		}
	}	
	
	router := gin.Default()
	// Redirect root to Swagger docs
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("", controllers.GetBooks)
			books.GET("/:id", controllers.GetBook)
			books.POST("", controllers.CreateBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		v1.GET("/health", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/health")
		})
	}

	// Health check
	router.GET("/health", controllers.HealthCheck)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
} 