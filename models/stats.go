package models

import "time"

type UsageStats struct {
	ID        string    `json:"id"`
	Endpoint  string    `json:"endpoint"`
	Method    string    `json:"method"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
} 