package configs

import (
	"ASSIGNMENT2/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{}
	config.WebServer.Port = os.Getenv("PORT")

	config.Database.Host = os.Getenv("PG_HOST")
	config.Database.Name = os.Getenv("PG_NAME")
	config.Database.Password = os.Getenv("PG_PASSWORD")
	config.Database.Port = os.Getenv("PG_PORT")
	config.Database.User = os.Getenv("PG_USER")

	return config
}
