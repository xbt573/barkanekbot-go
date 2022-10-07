// Bot command handlers
package handlers

import "gopkg.in/telebot.v3"
import "github.com/xbt573/barkanekbot-go/database"

// Handler for /anek command
func AnekHandler(ctx telebot.Context) error {
	anek, err := database.GetRandomAnek()
	if err != nil {
		return ctx.Reply("Внутренняя ошибка")
	}

	if anek == "" {
		return ctx.Reply("Анекдот пуст, возможно вы забыли заполнить базу")
	}

	return ctx.Reply(anek)
}
