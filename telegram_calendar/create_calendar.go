package telegram_calendar

import (
	"github.com/dykov/gocalendar"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"time"
)

type TelegramCalendar struct {
	bot               *tgbotapi.BotAPI
	firstDayOfTheWeek time.Weekday
}

func NewTelegramCalendar(bot *tgbotapi.BotAPI, weekStartsWith time.Weekday) TelegramCalendar {
	return TelegramCalendar{bot, weekStartsWith}
}

// Use year=0 to select the current year, and month=0 to select the current month.
func (tc *TelegramCalendar) CreateCalendar(year int, month time.Month) tgbotapi.InlineKeyboardMarkup {

	checkYearAndMonth(&year, &month)

	dataIgnore := createCallbackData("IGNORE", year, month, 0)

	var keyboard tgbotapi.InlineKeyboardMarkup
	var row []tgbotapi.InlineKeyboardButton

	row = append(row, tgbotapi.InlineKeyboardButton{
		Text: month.String() + " , " + strconv.Itoa(year), CallbackData: &dataIgnore,
	})
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	row = []tgbotapi.InlineKeyboardButton{}
	daysOfWeek := []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	for _, day := range daysOfWeek {
		row = append(row, tgbotapi.InlineKeyboardButton{Text: day, CallbackData: &dataIgnore})
	}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	monthCalendar := gocalendar.MonthCalendar(year, month, tc.firstDayOfTheWeek)

	for _, week := range monthCalendar {
		row = []tgbotapi.InlineKeyboardButton{}
		for _, day := range week {
			if day == -1 {
				row = append(row, tgbotapi.InlineKeyboardButton{Text: " ", CallbackData: &dataIgnore})
			} else {
				dayCallbackData := createCallbackData("DAY", year, month, day)
				row = append(row, tgbotapi.InlineKeyboardButton{Text: strconv.Itoa(day), CallbackData: &dayCallbackData})
			}
		}
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	row = []tgbotapi.InlineKeyboardButton{}
	prevMonthCallbackData := createCallbackData("PREV-MONTH", year, month, 0)
	nextMonthCallbackData := createCallbackData("NEXT-MONTH", year, month, 0)
	row = append(row, tgbotapi.InlineKeyboardButton{Text: "<", CallbackData: &prevMonthCallbackData})
	row = append(row, tgbotapi.InlineKeyboardButton{Text: " ", CallbackData: &dataIgnore})
	row = append(row, tgbotapi.InlineKeyboardButton{Text: ">", CallbackData: &nextMonthCallbackData})
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	return keyboard

}
