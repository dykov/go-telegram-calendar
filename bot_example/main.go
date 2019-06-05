package main

import (
	. "../telegram_calendar"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"time"
)

var myBot *tgbotapi.BotAPI
var updChannel tgbotapi.UpdatesChannel

func main() {

	AuthBot()

	tc := TelegramCalendar{}

	for update := range updChannel {

		switch {

		case update.Message != nil && update.Message.Text == `/calendar`:
			tc = NewTelegramCalendar(myBot, time.Monday)
			fmt.Println(tc.OpenCalendar(&update, ""))

		case update.CallbackQuery != nil:
			fmt.Println(tc.ProcessCalendarSelection(&update))

		case update.Message != nil && update.Message.Text != "":
			fmt.Println(tc.OpenCalendar(&update, ""))

		}

	}

}

func AuthBot() {

	var err error
	myBot, err = tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Fatalln(err)
	}

	myBot.Debug = false
	log.Infof("Authorized on account %s", myBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updChannel, err = myBot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatalln(err)
	}

}
