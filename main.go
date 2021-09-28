package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"support-bot/internal/handlers"
	"support-bot/internal/infrastructure/auth"
	"support-bot/internal/infrastructure/env"
	"support-bot/internal/infrastructure/logs"
	"support-bot/internal/persistence"
	"support-bot/internal/repository/telegram"
	"support-bot/internal/repository/viber"
	"support-bot/internal/server"
	"support-bot/internal/service"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)

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
	svc := service.New(repo, telegram.NewClient(), viber.NewClient(), authenticator)
	handler := handlers.NewHandler(svc, authenticator)

	srv := server.New(cfg.GetAddr(), cfg.Domain, handler)
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
