package process

import (
	"go_ensemble_bot/internal/ui/keyboard"
	"go_ensemble_bot/util"
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

func Do(b *telebot.Bot) {
	b.Handle("/card_number", func(c tele.Context) error {
		return c.Send("5228 6005 9599 3934")
	})

	b.Handle("/yandex_disc", func(c tele.Context) error {
		return c.Send("https://disk.yandex.ru/d/2OuuGeT8TxJu_w")
	})

	b.Handle("/nevesomo", func(c tele.Context) error {
		return c.Send("К. Меладзе - Невесомо", keyboard.NevesomoInlineKeyboard(b))
	})

	b.Handle("/rain", func(c tele.Context) error {
		return c.Send(util.SongTextMessage("Rain"))
	})

}
