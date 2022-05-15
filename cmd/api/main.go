package main

import (
	"context"
	"github.com/cnblvr/testing_forms/pkg/config"
	handlerHttp "github.com/cnblvr/testing_forms/pkg/handler/http"
	"github.com/cnblvr/testing_forms/pkg/logger"
	serverHttp "github.com/cnblvr/testing_forms/pkg/server/http"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create config")
	}
	logger.Init(cfg)

	handler := handlerHttp.New()
	handler.PrintHandlers()

	server := serverHttp.New(cfg, handler)
	go func() {
		if err := server.Run(); err != nil {
			log.Error().Err(err).Msg("error occurred while running http server")
		}
	}()
	log.Info().Msg("server started...")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("server shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("error occurred on server shutting down")
	}
}
