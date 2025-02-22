package utils

import "github.com/google/uuid"

// GenerateUUID creates a new UUID string
func GenerateUUID() string {
	return uuid.New().String()
}

// ValidateUUID checks if a string is a valid UUID
func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
} 