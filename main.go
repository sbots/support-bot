package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"support-bot/infrastructure/auth"
	"support-bot/infrastructure/env"
	"support-bot/logs"
	"support-bot/persistence"
	"support-bot/repository/telegram"
	"support-bot/repository/viber"
	"support-bot/server"
	"syscall"
)

func main() {
	cfg, err := env.FromOS()
	if err != nil {
		log.Fatal(err)
	}

	logs.InitLogger(cfg.LogLevel, cfg.LogPrettify)

	repo, err := persistence.NewRepository(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	authenticator, err := auth.NewAuthenticator(cfg.SecretKey, cfg.TokenIssuer, cfg.TokenExpiration)
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(cfg.GetAddr(), cfg.Domain, telegram.NewClient(), viber.NewClient(), repo, authenticator)
	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)

	srv.Run(ctx)
}

func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Println("Got Interrupt signal")
		stop()
	}()
}
