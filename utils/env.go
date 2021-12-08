package utils

import (
	"log"
	"os"
)

var (
	HOST      string
	PORT      string
	HTTPS     string
	PROXY     bool
	URL       string
	REDIS_URI string
)

func Init_env_vars() {

	UHOST, err := os.LookupEnv("HOST")
	if !err {
		log.Fatal("HOST enviroment variable not found, please set it!")
	}
	HOST = UHOST

	UHTTPS, _ := os.LookupEnv("HTTPS")
	if UHTTPS != "true" {
		HTTPS = "http"
	} else {
		HTTPS = "https"
	}

	UPROXY, _ := os.LookupEnv("PROXY")
	if UPROXY != "true" {
		PROXY = false
	} else {
		PROXY = true
	}

	UPORT, err := os.LookupEnv("PORT")
	if !err {
		PORT = "3000"
	} else {
		PORT = UPORT
	}

	UREDIS_URI, _ := os.LookupEnv("REDIS_URI")
	if UREDIS_URI != "" {
		REDIS_URI = "redis"
	}

	create_string()

}

func create_string() {
	if !PROXY {
		URL = HTTPS + "://" + HOST + ":" + PORT + "/"
	} else {
		URL = HTTPS + "://" + HOST + "/"
	}
}
