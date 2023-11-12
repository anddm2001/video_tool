package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config AppConfig

type AppConfig struct {
	MaxGoroutines int
	OutDir        string
	InDir         string
	FFMPEGPath    string
}

func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &AppConfig{
		MaxGoroutines: getEnvAsInt("MAX_GOROUTINES", 5),
		OutDir:        getEnv("OUT_DIR", "out"),
		InDir:         getEnv("IN_DIR", "in"),
		FFMPEGPath:    getEnv("FFMPEG_PATH", "/usr/bin/ffmpeg"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
