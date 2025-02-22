package main

import (
	"fmt"
	"go-rest-api/controllers"
	_ "go-rest-api/docs"
	"log"
	"net/http"
	"os"

	"go-rest-api/db"
	"go-rest-api/metrics"
	"go-rest-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Activity API
// @version         1.0
// @description     An API for tracking device activities and usage statistics.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Initialize activity controller
	controllers.InitActivityController(db.OB)

	for _ , arg := range os.Args {
		if arg == "healthcheck" {
			fmt.Println("OK")
			os.Exit(0)
		}
	}	
	
	// Initialize Prometheus metrics
	metrics.Init()
	
	router := gin.Default()
	// Redirect root to Swagger docs
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	// Add Prometheus middleware to all routes
	router.Use(middleware.PrometheusMiddleware())

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		stats := v1.Group("/stats")
		{
			stats.POST("", controllers.CreateStats)
			stats.GET("", controllers.GetAllStats)
			stats.GET("/endpoints/:endpoint", controllers.GetStatsByEndpoint)
			stats.DELETE("/endpoints/:endpoint", controllers.DeleteStatsByEndpoint)
			stats.DELETE("/:id", controllers.DeleteStats)
		}
		activities := v1.Group("/activities")
		{
			activities.POST("", controllers.CreateActivity)
			activities.GET("", controllers.GetAllActivities)
			activities.GET("/device/:device", controllers.GetActivitiesByDevice)
			activities.GET("/grid/:grid", controllers.GetActivitiesByGrid)
			activities.DELETE("/:id", controllers.DeleteActivity)
		}
		v1.GET("/health", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/health")
		})
	}

	// Health check
	router.GET("/health", controllers.HealthCheck)

	// Add metrics endpoint
	router.GET("/metrics", metrics.PrometheusHandler())

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
} 