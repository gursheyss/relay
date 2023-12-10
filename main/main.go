package main

import (
	"log"

	"github.com/gursheyss/relay"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = relay.Init()
	if err != nil {
		log.Println(err)
	}
}
