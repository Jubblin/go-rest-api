package controllers

import (
	"go-rest-api/metrics"
	"go-rest-api/models"
	"net/http"
	"time"

	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// In-memory storage for stats
var statsStore = make(map[string]models.UsageStats)

func init() {
	// Add sample stats
	sampleStats := []models.UsageStats{
		{
			ID:        uuid.New().String(),
			Endpoint:  "/api/v1/books",
			Method:    "GET",
			Status:    200,
			Timestamp: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        uuid.New().String(),
			Endpoint:  "/api/v1/books",
			Method:    "POST",
			Status:    201,
			Timestamp: time.Now().Add(-12 * time.Hour),
		},
		{
			ID:        uuid.New().String(),
			Endpoint:  "/health",
			Method:    "GET",
			Status:    200,
			Timestamp: time.Now().Add(-1 * time.Hour),
		},
	}

	for _, stat := range sampleStats {
		statsStore[stat.ID] = stat
	}
}

// CreateStats godoc
// @Summary Create usage statistics
// @Description Records new usage statistics
// @Tags stats
// @Accept json
// @Produce json
// @Param stats body models.UsageStats true "Stats Data"
// @Success 201 {object} models.UsageStats
// @Failure 400 {object} map[string]string
// @Router /stats [post]
func CreateStats(c *gin.Context) {
	defer metrics.StatsOperationsTotal.WithLabelValues("create").Inc()
	var newStats models.UsageStats
	if err := c.BindJSON(&newStats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStats.ID = utils.GenerateUUID()
	newStats.Timestamp = time.Now()
	statsStore[newStats.ID] = newStats
	c.JSON(http.StatusCreated, newStats)
}

// GetAllStats godoc
// @Summary Get all statistics
// @Description Retrieves all usage statistics
// @Tags stats
// @Produce json
// @Success 200 {array} models.UsageStats
// @Router /stats [get]
func GetAllStats(c *gin.Context) {
	defer metrics.StatsOperationsTotal.WithLabelValues("list").Inc()
	stats := make([]models.UsageStats, 0, len(statsStore))
	for _, stat := range statsStore {
		stats = append(stats, stat)
	}
	c.JSON(http.StatusOK, stats)
}

// GetStatsByEndpoint godoc
// @Summary Get statistics by endpoint
// @Description Retrieves statistics for a specific endpoint
// @Tags stats
// @Produce json
// @Param endpoint path string true "Endpoint Path"
// @Success 200 {array} models.UsageStats
// @Router /stats/endpoints/{endpoint} [get]
func GetStatsByEndpoint(c *gin.Context) {
	defer metrics.StatsOperationsTotal.WithLabelValues("get_by_endpoint").Inc()
	endpoint := c.Param("endpoint")
	filteredStats := make([]models.UsageStats, 0)

	for _, stat := range statsStore {
		if stat.Endpoint == endpoint {
			filteredStats = append(filteredStats, stat)
		}
	}
	c.JSON(http.StatusOK, filteredStats)
}

// DeleteStatsByEndpoint godoc
// @Summary Delete statistics by endpoint
// @Description Deletes all statistics for a specific endpoint
// @Tags stats
// @Produce json
// @Param endpoint path string true "Endpoint Path"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /stats/endpoints/{endpoint} [delete]
func DeleteStatsByEndpoint(c *gin.Context) {
	defer metrics.StatsOperationsTotal.WithLabelValues("delete_by_endpoint").Inc()
	endpoint := c.Param("endpoint")
	deleted := false

	for id, stat := range statsStore {
		if stat.Endpoint == endpoint {
			delete(statsStore, id)
			deleted = true
		}
	}

	if deleted {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No stats found for endpoint"})
	}
}

// DeleteStats godoc
// @Summary Delete statistics
// @Description Deletes specific statistics by ID
// @Tags stats
// @Produce json
// @Param id path string true "Stats ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /stats/{id} [delete]
func DeleteStats(c *gin.Context) {
	defer metrics.StatsOperationsTotal.WithLabelValues("delete").Inc()
	id := c.Param("id")

	if !utils.ValidateUUID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if _, exists := statsStore[id]; exists {
		delete(statsStore, id)
		c.Status(http.StatusNoContent)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Stats not found"})
} 