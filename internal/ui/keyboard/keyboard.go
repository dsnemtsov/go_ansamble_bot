package keyboard

import (
	"go_ensemble_bot/util"
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

func NevesomoInlineKeyboard(b *telebot.Bot) *tele.ReplyMarkup {
	selector := &tele.ReplyMarkup{}

	btnText := selector.Data("Текст", "text", "")
	btnAudio := selector.Data("Youtube", "audio", "")
	btnTracks := selector.Data("Партии", "tracks", "")

	btnAudio.URL = "https://youtu.be/gHjzgbxFoSw"

	selector.Inline(
		selector.Row(btnText, btnAudio, btnTracks),
	)

	b.Handle(&btnText, func(c tele.Context) error {
		return c.Send(util.SongTextMessage("Nevesomo"))
	})

	return selector
}
