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
	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}
	addr := "0.0.0.0:" + port

	srv := server.New(addr)
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
