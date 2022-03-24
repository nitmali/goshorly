package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HOST      string
	PORT      string
	HTTPS     string
	PROXY     bool
	URL       string
	REDIS_URI string
	ESTATS    string
)

func Init_env_file() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("INFO: ", "Error loading .env file using environment variables")
	}
}

func Init_env_vars() {
	fmt.Println("-- Initializing environment variables... --")

	UHOST, err := os.LookupEnv("HOST")
	if !err {
		log.Fatal("HOST enviroment variable not found, please set it!")
	}
	HOST = UHOST
	fmt.Println("HOST: ", HOST)

	UHTTPS, _ := os.LookupEnv("HTTPS")
	if UHTTPS != "true" {
		HTTPS = "http"
	} else {
		HTTPS = "https"
	}
	fmt.Println("Proto: ", HTTPS)

	UPROXY, _ := os.LookupEnv("PROXY")
	if UPROXY != "true" {
		PROXY = false
	} else {
		PROXY = true
	}
	fmt.Println("Own reverse proxy: ", PROXY)

	UPORT, err := os.LookupEnv("PORT")
	if !err {
		PORT = "3000"
	} else {
		PORT = UPORT
	}
	fmt.Println("Port: ", PORT)

	UREDIS_URI, err := os.LookupEnv("REDIS_URI")
	if !err {
		REDIS_URI = "redis:6379"
	} else {
		REDIS_URI = UREDIS_URI
	}
	fmt.Println("Redis URI: ", REDIS_URI)

	UESTATS, _ := os.LookupEnv("ENABLE_STATS")
	if UESTATS != "false" {
		ESTATS = "true"
	} else {
		ESTATS = "false"
	}
	fmt.Println("Stats enabled: ", ESTATS)

	create_string()

}

func create_string() {
	if !PROXY {
		URL = HTTPS + "://" + HOST + ":" + PORT + "/"
	} else {
		URL = HTTPS + "://" + HOST + "/"
	}

	fmt.Println("URL: ", URL)
	fmt.Println("-- Environment variables initialized! / Starting Webserver --")
}
