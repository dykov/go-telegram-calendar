package telegram_calendar

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"time"
)

func (tc *TelegramCalendar) ProcessCalendarSelection(update *tgbotapi.Update) (time.Time, error) {

	var separatedCallbackData = separateCallbackData(update.CallbackQuery.Data)
	var action = separatedCallbackData[0]
	var year, _ = strconv.ParseInt(separatedCallbackData[1], 10, 64)
	var month, _ = strconv.ParseInt(separatedCallbackData[2], 10, 64)
	var day, _ = strconv.ParseInt(separatedCallbackData[3], 10, 64)

	var selectedMonth = time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, time.Local)

	switch action {

	case "IGNORE":
		_, _ = tc.bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{CallbackQueryID: update.CallbackQuery.ID})

	case "DAY":
		return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local), nil

	case "PREV-MONTH":
		prevMonth := selectedMonth.AddDate(0, -1, 0)
		var msg = tgbotapi.NewEditMessageText(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			update.CallbackQuery.Message.Text,
		)
		var msgMarkup = tgbotapi.NewEditMessageReplyMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tc.CreateCalendar(prevMonth.Year(), prevMonth.Month()),
		)
		msg.ReplyMarkup = msgMarkup.ReplyMarkup
		if _, err := tc.bot.Send(msg); err != nil {
			return time.Time{}, err
		}

	case "NEXT-MONTH":
		nextMonth := selectedMonth.AddDate(0, 1, 0)
		var msg = tgbotapi.NewEditMessageText(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			update.CallbackQuery.Message.Text,
		)
		var msgMarkup = tgbotapi.NewEditMessageReplyMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tc.CreateCalendar(nextMonth.Year(), nextMonth.Month()),
		)
		msg.ReplyMarkup = msgMarkup.ReplyMarkup
		if _, err := tc.bot.Send(msg); err != nil {
			return time.Time{}, err
		}

	}

	return time.Time{}, nil
}
