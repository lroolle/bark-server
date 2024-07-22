package database

import (
	"fmt"
	"os"

	"github.com/mritd/logger"
)

// EnvBase represents a database that uses environment variables for storage
type EnvBase struct {
	key         string
	deviceToken string
}

// NewEnvBase creates a new EnvBase instance and validates the required environment variables
func NewEnvBase() Database {
	key := os.Getenv("BARK_KEY")
	deviceToken := os.Getenv("BARK_DEVICE_TOKEN")

	if key == "" || deviceToken == "" {
		logger.Fatalf("BARK_KEY and BARK_DEVICE_TOKEN must be set in serverless mode.")
	}

	return &EnvBase{key: key, deviceToken: deviceToken}
}

// CountAll returns the count of all entries in the database
func (d *EnvBase) CountAll() (int, error) {
	return 1, nil
}

// DeviceTokenByKey returns the device token associated with the given key
func (d *EnvBase) DeviceTokenByKey(key string) (string, error) {
	if key == d.key {
		return d.deviceToken, nil
	}
	return "", fmt.Errorf("key not found")
}

// SaveDeviceTokenByKey saves the device token associated with the given key
func (d *EnvBase) SaveDeviceTokenByKey(key, token string) (string, error) {
	if token == d.deviceToken {
		return d.key, nil
	}
	return "", fmt.Errorf("device token is invalid")
}

// Close closes the database connection (no-op for EnvBase)
func (d *EnvBase) Close() error {
	return nil
}
