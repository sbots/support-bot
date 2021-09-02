package main

import (
	"context"
	"log"
	auth2 "support-bot/infrastructure/auth"
	env2 "support-bot/infrastructure/env"
	persistence2 "support-bot/persistence"
	telegram2 "support-bot/repository/telegram"
	viber2 "support-bot/repository/viber"
	server2 "support-bot/server"
)

func main() {
	cfg, err := env2.FromOS()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := persistence2.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	authenticator, err := auth2.NewAuthenticator(cfg.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	srv := server2.New(cfg.GetAddr(), cfg.Domain, telegram2.NewClient(), viber2.NewClient(), repo, authenticator, distFS)

	ctx := context.Background()
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
