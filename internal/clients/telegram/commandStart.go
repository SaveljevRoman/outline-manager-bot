package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tgc *TgBotClient) CommandStart() ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		tgc.delMsgNoErr(update.FromChat().ID, update.Message.MessageID)

		_, err := tgc.Storage.PgClient.InsertNewOwner(ctx, update.FromChat().ID, update.Message.From.UserName)
		if err != nil {
			fmt.Println(err)
		}

		// Если пользователь уже существует в базе, то просто выходим и ничего не пишем
		//if !owner.IsInserted {
		//	return nil
		//}

		msg := tgc.mainMessage(update.FromChat().ID)

		if _, err = bot.Send(msg); err != nil {
			return err
		}
		return nil
	}
}

func (tgc *TgBotClient) CommandCancelStart() ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		var mesId int
		if update.Message != nil {
			mesId = update.Message.MessageID
		}

		if mesId == 0 && update.CallbackQuery != nil {
			mesId = update.CallbackQuery.Message.MessageID
		}

		tgc.delMsgNoErr(update.FromChat().ID, mesId)

		msg := tgc.mainMessage(update.FromChat().ID)
		if _, err := bot.Send(msg); err != nil {
			return err
		}
		return nil
	}
}

func (tgc *TgBotClient) mainMessage(chatId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		chatId,
		"Привет, этот бот создан для удообного управления ключами вашего outline сервера.",
	)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
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

	return msg
}
