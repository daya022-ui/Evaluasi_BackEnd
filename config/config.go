package config

import (
	"fmt"
	"perpustakaan/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfigurationMap struct {
	Port         int
	IsProduction bool
	DbUri        string
}

var config *AppConfigurationMap

func Get() *AppConfigurationMap {
	return config
}

func Load() {
	log.Println("load config from environment")
	// Get the configuration
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables, try to get from environtment OS")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	// set default value port if env doesn't have PORT config
	if err != nil {
		port = 8080
	}

	isProduction := utils.SafeCompareString(os.Getenv("IS_PRODUCTION"), "true")

	// set global variable config
	config = &AppConfigurationMap{
		Port:         port,
		IsProduction: isProduction,
		DbUri:        loadDatabaseConfig(),
	}
}

func loadDatabaseConfig() string {
	user := getFromEnv("DB_USER")
	pass := getFromEnv("DB_PASS")
	name := getFromEnv("DB_NAME")
	host := getFromEnv("DB_HOST")
	port := getFromEnv("DB_PORT")
	timeZone := getFromEnv("DB_TIME_ZONE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", host, user, pass, name, port, timeZone)
}

func getFromEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is not set", value)
	}

	return value
}
