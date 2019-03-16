package main

import (
	"github.com/joho/godotenv"
	"log"
	"lottomonitor/eurojackpot"
)

func main() {
	loadEnv()

	eurojackpot.CheckNumbers()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}