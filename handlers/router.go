// Bot command handlers
package handlers

import "gopkg.in/telebot.v3"

// Route commands
func Router(bot *telebot.Bot) {
	bot.Handle("/start", StartHandler)
	bot.Handle("/help", HelpHandler)
	bot.Handle("/anek", AnekHandler)
}
