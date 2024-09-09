package telegram

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"outline-manager-bot/config"
	"outline-manager-bot/internal/storage"
	"runtime/debug"
	"time"
)

const logPointTgClientUpdate = "point: tg client update "

type TgBotClient struct {
	client        *tgbotapi.BotAPI
	updateTimeout int
	updateOffset  int
	debug         bool
	cmdViews      map[string]ViewFunc
	Storage       *storage.Storage
}

type ViewFunc func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error

func NewTgBotClient(ctx context.Context, tgConf *config.TGConfig, pgConf *config.PostgresConfig) (*TgBotClient, error) {
	if tgConf.TgBotToken == "" {
		return nil, errors.New("telegram bot token not found")
	}

	client, err := tgbotapi.NewBotAPI(tgConf.TgBotToken)
	if err != nil {
		return nil, err
	}

	return &TgBotClient{
		client:        client,
		debug:         tgConf.Debug,
		updateOffset:  tgConf.UpdateOffset,
		updateTimeout: tgConf.UpdateTimeout,
		Storage:       storage.NewStorage(ctx, pgConf),
	}, nil
}

func (tgc *TgBotClient) RegisterCmdView(cmd string, view ViewFunc) {
	if tgc.cmdViews == nil {
		tgc.cmdViews = make(map[string]ViewFunc)
	}

	tgc.cmdViews[cmd] = view
}

func (tgc *TgBotClient) Run(ctx context.Context) error {
	tgc.client.Debug = tgc.debug

	u := tgbotapi.NewUpdate(tgc.updateOffset)
	u.Timeout = tgc.updateTimeout

	updates := tgc.client.GetUpdatesChan(u)
	for {
		select {
		case update := <-updates:
			updateCtx, updateCancel := context.WithTimeout(ctx, 5*time.Second)
			tgc.handleUpdate(updateCtx, update)
			updateCancel()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (tgc *TgBotClient) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("%s. panic recovered: %v\n%s", logPointTgClientUpdate, p, string(debug.Stack()))
		}
	}()

	if update.Message == nil || !update.Message.IsCommand() {
		return
	}

	if !update.Message.IsCommand() {
		return
	}

	cmd := update.Message.Command()

	cmdView, ok := tgc.cmdViews[cmd]
	if !ok {
		return
	}

	if err := cmdView(ctx, tgc.client, update); err != nil {
		log.Printf("%s. failed to handle update: %v", logPointTgClientUpdate, err)

		if _, err := tgc.client.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "internal error")); err != nil {
			log.Printf("%s failed to send message: %v", logPointTgClientUpdate, err)
		}
	}
}
