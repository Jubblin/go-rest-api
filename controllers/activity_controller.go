package controllers

import (
	"go-rest-api/metrics"
	"go-rest-api/models"
	"net/http"
	"time"

	"go-rest-api/repositories"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/objectbox/objectbox-go/objectbox"
)

type ActivityController struct {
	repo *repositories.ActivityRepository
}

var activityController ActivityController

// InitActivityController initializes the controller after DB setup
func InitActivityController(ob *objectbox.ObjectBox) {
	activityController = ActivityController{
		repo: repositories.NewActivityRepository(ob),
	}

	// Sample activity data
	sampleActivities := []models.DeviceActivity{
		{
			UniqueId:   utils.GenerateUUID(),
			SourceIP:   "192.168.1.100",
			DeviceName: "device-alpha",
			GridName:   "grid-east",
			Action:     "login",
			Timestamp:  time.Now().Add(-1 * time.Hour),
		},
		{
			UniqueId:   utils.GenerateUUID(),
			SourceIP:   "192.168.1.101",
			DeviceName: "device-beta",
			GridName:   "grid-west",
			Action:     "data_sync",
			Timestamp:  time.Now().Add(-30 * time.Minute),
		},
	}

	for _, activity := range sampleActivities {
		activityController.repo.Create(activity)
	}

	updateActivityMetrics()
}

// Helper function to update Prometheus metrics
func updateActivityMetrics() {
	// Reset all metrics
	metrics.ActivityCount.Reset()

	// Get all activities from repository
	activities, err := activityController.repo.GetAll()
	if err != nil {
		return
	}

	// Count activities by grid and device
	gridCounts := make(map[string]int)
	deviceCounts := make(map[string]int)
	
	for _, activity := range activities {
		metrics.ActivityCount.WithLabelValues(activity.GridName, activity.DeviceName).Inc()
		gridCounts[activity.GridName]++
		deviceCounts[activity.DeviceName]++
	}
}

// generateUUID creates a new UUID string
func generateUUID() string {
	return uuid.New().String()
}

// validateUUID checks if a string is a valid UUID
func validateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// CreateActivity godoc
// @Summary Create a new activity
// @Description Records a new device activity with headers
// @Tags activities
// @Accept json
// @Produce json
// @Param activity body models.DeviceActivity true "Activity Data"
// @Success 201 {object} models.DeviceActivity
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /activities [post]
func CreateActivity(c *gin.Context) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.ActivityLatency.WithLabelValues("create").Observe(duration)
	}()

	var newActivity models.DeviceActivity
	if err := c.BindJSON(&newActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	headers := make(map[string]string)
	for key, values := range c.Request.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}
	if err := newActivity.SetHeaders(headers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newActivity.UniqueId = utils.GenerateUUID()
	newActivity.Timestamp = time.Now()

	err := activityController.repo.Create(newActivity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.ActivityOperationsTotal.WithLabelValues("create", newActivity.GridName, newActivity.DeviceName).Inc()
	c.JSON(http.StatusCreated, newActivity)
}

// GetAllActivities godoc
// @Summary Get all activities
// @Description Retrieves all recorded device activities
// @Tags activities
// @Produce json
// @Success 200 {array} models.DeviceActivity
// @Failure 500 {object} map[string]string
// @Router /activities [get]
func GetAllActivities(c *gin.Context) {
	activities, err := activityController.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// GetActivitiesByDevice godoc
// @Summary Get activities by device
// @Description Retrieves activities for a specific device
// @Tags activities
// @Produce json
// @Param device path string true "Device Name"
// @Success 200 {array} models.DeviceActivity
// @Failure 500 {object} map[string]string
// @Router /activities/device/{device} [get]
func GetActivitiesByDevice(c *gin.Context) {
	deviceName := c.Param("device")
	activities, err := activityController.repo.GetByDevice(deviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// DeleteActivity godoc
// @Summary Delete an activity
// @Description Deletes a specific activity by ID
// @Tags activities
// @Produce json
// @Param id path string true "Activity ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /activities/{id} [delete]
func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	if !utils.ValidateUUID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	err := activityController.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetActivitiesByGrid godoc
// @Summary Get activities by grid
// @Description Retrieves activities for a specific grid
// @Tags activities
// @Produce json
// @Param grid path string true "Grid Name"
// @Success 200 {array} models.DeviceActivity
// @Failure 500 {object} map[string]string
// @Router /activities/grid/{grid} [get]
func GetActivitiesByGrid(c *gin.Context) {
	gridName := c.Param("grid")
	activities, err := activityController.repo.GetByGrid(gridName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
} 