package main

import (
	"context"
	"log"
	"support-bot/api/auth"
	"support-bot/api/config"
	"support-bot/api/persistence"
	"support-bot/api/repository/telegram"
	"support-bot/api/repository/viber"
	"support-bot/api/server"
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
