// Bot command handlers
package handlers

import "gopkg.in/telebot.v3"

// Handler for /start command
func StartHandler(ctx telebot.Context) error {
	return ctx.Reply(`
Привет! Этот бот собирает анекдоты с различных групп и отправляет их в вашу любимую беседу
Чтобы получить рандомный анекдот напишите команду /anek, права администратора боту не нужны :).`)
}
