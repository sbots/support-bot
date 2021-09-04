package main

import (
	"context"
	"log"
	"support-bot/infrastructure/auth"
	"support-bot/infrastructure/env"
	"support-bot/persistence"
	"support-bot/repository/telegram"
	"support-bot/repository/viber"
	"support-bot/server"
)

func main() {
	cfg, err := env.FromOS()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := persistence.NewRepository(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	authenticator, err := auth.NewAuthenticator(cfg.SecretKey, cfg.TokenIssuer, cfg.TokenExpiration)
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(cfg.GetAddr(), cfg.Domain, telegram.NewClient(), viber.NewClient(), repo, authenticator)

	ctx := context.Background()
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
