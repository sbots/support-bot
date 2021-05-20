package main

import (
	"context"
	"log"
	"support-bot/server"
)

func main() {
	ctx := context.Background()
	srv := server.New("localhost:7777")
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
