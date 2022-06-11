package process

import (
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

func Do(b *telebot.Bot) {
	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/card_number", func(c tele.Context) error {
		return c.Send("5228 6005 9599 3934")
	})
}
