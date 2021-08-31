package main

import (
	"context"
	"log"
	auth2 "support-bot/service/auth"
	config2 "support-bot/service/config"
	persistence2 "support-bot/service/persistence"
	telegram2 "support-bot/service/repository/telegram"
	viber2 "support-bot/service/repository/viber"
	server2 "support-bot/service/server"
)

func main() {
	ctx := context.Background()

	cfg, err := config2.FromOS()
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
	srv := server2.New(cfg.GetAddr(), cfg.Domain, telegram2.NewClient(), viber2.NewClient(), repo, authenticator)

	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
