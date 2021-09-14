package main

import (
	"backup-service/app"
	backup2 "backup-service/backup"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
)

var rdb *redis.Client
var ctx = context.Background()

func main() {
	log.Println("Starting backup service...")
	check(connectToRedis())

	backup := backup2.NewBackup(rdb)
	go backup.SubscribeToExpiredEvents(ctx)

	application := app.New()
	http.HandleFunc("/", application.Router.ServeHTTP)

	log.Println("Application started at port 8000")
	err := http.ListenAndServe(":8000", nil)
	check(err)
}

func connectToRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	return err
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
