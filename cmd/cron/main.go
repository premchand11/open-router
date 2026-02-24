package main

import (
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "redis:6379"},
		asynq.Config{
			Concurrency: 10,
		},
	)

	log.Println("Worker started...")

	if err := srv.Run(asynq.NewServeMux()); err != nil {
		log.Fatal(err)
	}
}