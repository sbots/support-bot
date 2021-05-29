package main

import (
	"context"
	"log"
	"support-bot/bots/telegram"
	"support-bot/config"
	"support-bot/server"
)

const testBotID = "1"

func main() {
	ctx := context.Background()

	cfg, err := config.FromOS()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(cfg.GetAddr())

	path := cfg.Domain + srv.GetEndpointForBot(testBotID)
	if err := telegram.ConnectNewBot(cfg.TestToken, path); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
