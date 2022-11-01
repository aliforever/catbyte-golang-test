package main

import (
	reporting_api "catbyte-golang-test/lib/reporting-api"
	"catbyte-golang-test/shared/logger"
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "",
		DB:       0,
	})

	// This is to make sure redis is alive
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	err = reporting_api.NewReportingAPIService(rdb, logger.DefaultLogger{}).Start(":8081")
	if err != nil {
		panic(err)
	}
}
