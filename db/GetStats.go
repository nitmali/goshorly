package db

import (
	"log"
	"strconv"
)

func GetTotalLinks() int {
	val1, err1 := Client.Get("total-links").Result()
	if err1 != nil {
		log.Fatal(err1)
	}
	val2, err2 := strconv.Atoi(val1)
	if err2 != nil {
		log.Fatal(err1)
	}
	return val2
}

func GetTotalViews() int {
	val1, err1 := Client.Get("total-views").Result()
	if err1 != nil {
		log.Fatal(err1)
	}
	val2, err2 := strconv.Atoi(val1)
	if err2 != nil {
		log.Fatal(err1)
	}
	return val2
}
