package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env struct for env file
type Env struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
	SSLMode  string
}

// LoadEnv return env file to struct
func LoadEnv() (c Env) {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	//Env Postgre
	c.Database = os.Getenv("DATABASE")
	c.User = os.Getenv("USERNAME")
	c.Password = os.Getenv("PASSWORD")
	c.Host = os.Getenv("HOST")
	c.Port = os.Getenv("PORT")
	c.SSLMode = os.Getenv("PORT")

	//Env Cloudinary

	return c
}
