// Bot command handlers
package handlers

import "gopkg.in/telebot.v3"

// Handler for /help command
func HelpHandler(ctx telebot.Context) error {
	return ctx.Reply(`
Команда у бота всего одна: /anek, права администратора для отправки анеков не требуются :).`)
}
