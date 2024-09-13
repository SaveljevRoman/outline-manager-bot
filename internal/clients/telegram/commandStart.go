package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tgc *TgBotClient) CommandStart() ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		delMsg := tgbotapi.NewDeleteMessage(update.FromChat().ID, update.Message.MessageID)
		if _, err := bot.Request(delMsg); err != nil {
			return err
		}

		var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить сервер", "/add_outline_server"),
				tgbotapi.NewInlineKeyboardButtonData("Ваш список серверов", "/show_outline_servers"),
			),
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("3", "/command_3"),
			//	tgbotapi.NewInlineKeyboardButtonData("4", "/command_4"),
			//	tgbotapi.NewInlineKeyboardButtonData("5", "/command_5"),
			//),
		)

		owner, err := tgc.Storage.PgClient.InsertNewOwner(ctx, update.FromChat().ID, update.Message.From.UserName)
		if err != nil {
			fmt.Println(err)
		}

		// Если пользователь уже существует в базе, то просто выходим и ничего не пишем
		if !owner.IsInserted {
			return nil
		}

		msg := tgbotapi.NewMessage(
			update.FromChat().ID,
			"Привет, этот бот создан для удообного управления ключами вашего outline сервера.",
		)
		msg.ReplyMarkup = numericKeyboard

		if _, err = bot.Send(msg); err != nil {
			return err
		}
		return nil
	}
}
