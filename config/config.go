package config

import (
	"os"
	"strconv"
)

type TwilioConfig struct {
	Accountsid string
	Authtoken  string
}

type Config struct {
	Cred     TwilioConfig
	sender   int
	receiver int
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Cred: TwilioConfig{
			Accountsid: getEnv("ACCOUNT_SID", ""),
			Authtoken:  getEnv("AUTH_TOKEN", ""),
		},
		receiver : getEnvAsInt("RECIPIENT_NUM", 1),
		sender : getEnvAsInt("TWIL_NUM", 1),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
	return value
    }

    return defaultVal
}