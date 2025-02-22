package repositories

import (
	"go-rest-api/metrics"
	"go-rest-api/models"
	"time"

	"github.com/objectbox/objectbox-go/objectbox"
)

type ActivityRepository struct {
	box *models.DeviceActivityBox
}

func NewActivityRepository(ob *objectbox.ObjectBox) *ActivityRepository {
	box := models.BoxForDeviceActivity(ob)
	repo := &ActivityRepository{box: box}
	repo.updateMetrics()
	return repo
}

func (r *ActivityRepository) updateMetrics() {
	count, err := r.box.Count()
	if err == nil {
		metrics.ObjectBoxEntityCount.WithLabelValues("activity").Set(float64(count))
	}
}

func (r *ActivityRepository) Create(activity models.DeviceActivity) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.ObjectBoxOperationDuration.WithLabelValues("create", "activity").Observe(duration)
	}()

	_, err := r.box.Put(&activity)
	if err != nil {
		return err
	}

	metrics.ObjectBoxOperationsTotal.WithLabelValues("create", "activity").Inc()
	r.updateMetrics()
	return nil
}

func (r *ActivityRepository) GetAll() ([]models.DeviceActivity, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.ObjectBoxOperationDuration.WithLabelValues("get_all", "activity").Observe(duration)
	}()

	results, err := r.box.GetAll()
	if err != nil {
		return nil, err
	}

	activities := make([]models.DeviceActivity, len(results))
	for i, result := range results {
		activities[i] = *result
	}

	metrics.ObjectBoxOperationsTotal.WithLabelValues("get_all", "activity").Inc()
	return activities, nil
}

func (r *ActivityRepository) GetByGrid(gridName string) ([]models.DeviceActivity, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.ObjectBoxOperationDuration.WithLabelValues("get_by_grid", "activity").Observe(duration)
	}()

	query := r.box.Query(models.DeviceActivity_.GridName.Equals(gridName, true))
	results, err := query.Find()
	if err != nil {
		return nil, err
	}

	activities := make([]models.DeviceActivity, len(results))
	for i, result := range results {
		activities[i] = *result
	}

	metrics.ObjectBoxOperationsTotal.WithLabelValues("get_by_grid", "activity").Inc()
	return activities, nil
}

func (r *ActivityRepository) GetByDevice(deviceName string) ([]models.DeviceActivity, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.ObjectBoxOperationDuration.WithLabelValues("get_by_device", "activity").Observe(duration)
	}()

	query := r.box.Query(models.DeviceActivity_.DeviceName.Equals(deviceName, true))
	results, err := query.Find()
	if err != nil {
		return nil, err
	}

	activities := make([]models.DeviceActivity, len(results))
	for i, result := range results {
		activities[i] = *result
	}

	metrics.ObjectBoxOperationsTotal.WithLabelValues("get_by_device", "activity").Inc()
	return activities, nil
}

func (r *ActivityRepository) Delete(uniqueId string) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.ObjectBoxOperationDuration.WithLabelValues("delete", "activity").Observe(duration)
	}()

	query := r.box.Query(models.DeviceActivity_.UniqueId.Equals(uniqueId, true))
	results, err := query.Find()
	if err != nil || len(results) == 0 {
		return err
	}

	err = r.box.Remove(results[0])
	if err != nil {
		return err
	}

	metrics.ObjectBoxOperationsTotal.WithLabelValues("delete", "activity").Inc()
	r.updateMetrics()
	return nil
} 