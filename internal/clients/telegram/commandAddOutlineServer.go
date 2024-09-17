package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	outlineInstall = "sudo bash -c \"$(wget -qO- https://raw.githubusercontent.com/Jigsaw-Code/outline-server/master/src/server_manager/install_scripts/install_server.sh)\""
	outlineExample = "\\{\\\"apiUrl\\\":\\\"https://xxx\\.xxx\\.xxx\\.xxx:xxxx/xxxxxxxxxxxxxxxxxxxxxx\\\",\\\"certSha256\\\":\\\"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\\\"\\}"
)

func (tgc *TgBotClient) CommandAddOutlineServer() ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		tgc.delMsgNoErr(update.FromChat().ID, update.CallbackQuery.Message.MessageID)

		msgText := fmt.Sprintf(
			"Для установки outline на Вашем сервере, выполните комманду:\n%s\n\n"+
				"Отправьте информацию, полученную после выполнения скрипта установки:\n %s",
			"```"+outlineInstall+"```",
			"`"+outlineExample+"`",
		)

		msg := tgbotapi.NewMessage(
			update.FromChat().ID,
			msgText,
		)
		msg.ParseMode = tgbotapi.ModeMarkdownV2
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("В начало", "/cancel_start"),
			),
		)

		if _, err := bot.Send(msg); err != nil {
			return err
		}

		return nil
	}
}
