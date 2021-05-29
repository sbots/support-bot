package main

import (
	"context"
	"log"
	"os"
	"support-bot/bots/telegram"
	"support-bot/server"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("bot")
	domain := os.Getenv("domain")

	srv := server.New("localhost:7777")
	path := domain + srv.GetEndpointForBot("1")
	bot, err := telegram.AddNewBot(token)
	if err != nil {
		log.Fatal(err)
	}
	if err := bot.SetNewWebhook(path); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
