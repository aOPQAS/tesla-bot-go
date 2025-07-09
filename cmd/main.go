package main

import (
	"os"
	"strconv"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"

	"github.com/tesla/tesla-bot-go/config"
	"github.com/tesla/tesla-bot-go/internal/deps"
	"github.com/tesla/tesla-bot-go/internal/pgsql"
	"github.com/tesla/tesla-bot-go/internal/server"
	tgnotifier "github.com/tesla/tesla-bot-go/internal/telegram"
	"github.com/tesla/tesla-bot-go/pkg/log"
	"go.uber.org/zap"
)

func main() {
	if err := os.Setenv("TZ", "UTC"); err != nil {
		log.Fatal("failed to set UTC timezone", zap.Error(err))
	}

	log.Info("loading .env file")
	if err := godotenv.Load(".env"); err != nil {
		log.Warn("failed to load .env file, using default settings")
	}

	log.Info("loading config")
	cfg := config.NewConfig()
	if err := env.Parse(cfg); err != nil {
		log.Fatal("failed to parse env config", zap.Error(err))
	}

	log.SetLogEncoding(cfg.Logger.Encoding)
	log.SetLogLevel("INFO")

	log.Info("logger initialized",
		zap.String("encoding", cfg.Logger.Encoding),
		zap.String("server_port", cfg.Server.Port),
	)

	log.Info("creating pgsql connection")
	conn, err := pgsql.CreatePostgresConnection(cfg.Postgres)
	if err != nil {
		log.Fatal("Failed to make pg connection", zap.Error(err))
	}

	log.Info("loading pgsql client")
	pg := pgsql.NewCLient(conn)

	log.Info("running migrations")
	if err := pg.RunMigrations(); err != nil {
		log.Fatal("failed to run migrations", zap.Error(err))
	}

	log.Info("starting server")

	myID, err := strconv.Atoi(cfg.Telegram.MyID)
	if err != nil {
		log.Fatal("Failed to convert my ID", zap.Error(err))
	}

	notifier, err := tgnotifier.New(cfg.Telegram.BotToken, int64(myID))
	if err != nil {
		log.Info("Failed to create Bot Instance", zap.Error(err))
	}

	s := server.New(&deps.Deps{
		PG: pg,
		TG: notifier,
	})

	if err := s.App.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
