package main

import (
	"context"
	"log"
	"support-bot/auth"
	"support-bot/config"
	"support-bot/persistence"
	"support-bot/repository/telegram"
	"support-bot/repository/viber"
	"support-bot/server"
)

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
	authenticator, err := auth.NewAuthenticator(cfg.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(cfg.GetAddr(), cfg.Domain, telegram.NewClient(), viber.NewClient(), repo, authenticator)

	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
