package app

import (
	"context"
	"errors"
	"log"
	"outline-manager-bot/config"
	"outline-manager-bot/internal/clients/telegram"
)

const logPointStart = "point: start application"

type App struct {
	Context  context.Context
	TgClient *telegram.TgBotClient
}

func NewApp(ctx context.Context, cfg *config.Config) *App {
	tgClient, err := telegram.NewTgBotClient(cfg.TgConf, cfg.PgConf)
	tgClient.RegisterCmdView("start", tgClient.CommandStart())

	if err != nil {
		log.Fatalf("%s. %v", logPointStart, err)
	}

	return &App{
		Context:  ctx,
		TgClient: tgClient,
	}
}

func (a *App) Start() {
	if err := a.TgClient.Run(a.Context); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Fatalf("%s. failed to run bot: %v", logPointStart, err)
			return
		}
	}
}
