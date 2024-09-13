package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// const PATH_API string = "http://10.104.0.14:7071/"

func GetApiPath() string {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	appsApi := os.Getenv("PATH_API")
	if appsApi == "" {
		appsApi = "http://174.138.22.4:7071/"
	}

	return appsApi
}
