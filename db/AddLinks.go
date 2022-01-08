package db

import (
	"strconv"
)

func StatsIncreaseTotalLinks() (int, error) {
	val, err := Client.Get("total-links").Result()

	if err != nil {
		err := Client.Set("total-links", "0", 0).Err()
		if err != nil {
			return 0, err
		}
	}
	i, _ := strconv.Atoi(val)
	i++
	err = Client.Set("total-links", i, 0).Err()
	if err != nil {
		return 0, err
	}
	return i, nil
}
