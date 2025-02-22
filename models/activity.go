package models

import (
	"encoding/json"
	"time"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type DeviceActivity struct {
	Id         uint64    `objectbox:"id"`
	UniqueId   string    `objectbox:"unique"`
	SourceIP   string
	DeviceName string    `objectbox:"index"`
	GridName   string    `objectbox:"index"`
	Action     string
	Headers    string    // Store as JSON string
	Timestamp  time.Time
}

// Helper methods for headers
func (a *DeviceActivity) SetHeaders(headers map[string]string) error {
	data, err := json.Marshal(headers)
	if err != nil {
		return err
	}
	a.Headers = string(data)
	return nil
}

func (a *DeviceActivity) GetHeaders() (map[string]string, error) {
	var headers map[string]string
	if err := json.Unmarshal([]byte(a.Headers), &headers); err != nil {
		return nil, err
	}
	return headers, nil
} 