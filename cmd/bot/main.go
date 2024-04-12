package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patytema/TinyClicker/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("couldn't load config: %w", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		return fmt.Errorf("couldn't start bot: %w", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var msg tgbotapi.MessageConfig
			if update.Message.Text == "/start" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "buttons")

			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "I didn't understand you")
			}

			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("could not reply: %s", err.Error())
			}
		}
	}

	return nil
}
