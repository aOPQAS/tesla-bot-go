package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramNotifier struct {
	bot  *tgbotapi.BotAPI
	MyID int64
}

// FUNC FOR CREATE TGBOT
func New(BotToken string, MyID int64) (*TelegramNotifier, error) {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		return nil, err
	}

	return &TelegramNotifier{
		bot:  bot,
		MyID: MyID,
	}, nil
}

// NOTIFICATION
func (tg *TelegramNotifier) Notify(data string) error {
	// create message
	msg := tgbotapi.NewMessage(tg.MyID, data)

	// send message
	_, err := tg.bot.Send(msg)
	return err
}
