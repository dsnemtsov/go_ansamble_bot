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

	b.Handle(&btnTracks, func(c tele.Context) error {
		return c.Send("Партии", nevesomoAudioInlineKeyboard(b))
	})

	return selector
}

func nevesomoAudioInlineKeyboard(b *telebot.Bot) *tele.ReplyMarkup {
	selector := &tele.ReplyMarkup{}

	btnMen1 := selector.Data("Муж. 1 куплет", "men1", "")
	btnWomen11 := selector.Data("Жен. 1-1 куплет", "women11", "")
	btnWomen12 := selector.Data("Жен. 1-2 куплет", "women12", "")

	b.Handle(&btnMen1, func(c tele.Context) error {
		return c.Send(util.AudioMessage("nevesomo", "men1"))
	})

	b.Handle(&btnWomen11, func(c tele.Context) error {
		return c.Send(util.AudioMessage("nevesomo", "women11"))
	})

	b.Handle(&btnWomen12, func(c tele.Context) error {
		return c.Send(util.AudioMessage("nevesomo", "women12"))
	})

	selector.Inline(
		selector.Row(btnMen1, btnWomen11, btnWomen12),
	)

	return selector
}
