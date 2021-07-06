package main

import (
	"context"
	"log"
	"support-bot/bots/telegram"
	"support-bot/config"
	"support-bot/persistence"
	"support-bot/server"
)

const testBotID = "1"

func main() {
	ctx := context.Background()

	cfg, err := config.FromOS()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := persistence.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(cfg.GetAddr(), cfg.Domain, telegram.NewClient(), repo)

	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
