package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Cek apakah file .env ada
	if _, err := os.Stat(".env"); err == nil {
		// Kalau ada, load
		err = godotenv.Load()
		if err != nil {
			log.Println("Env File not loaded")
		} else {
			log.Println("Env file loaded")
		}
	} else {
		// Kalau tidak ada, anggap ini mode Docker / production
		log.Println("env file not found")
	}
}
