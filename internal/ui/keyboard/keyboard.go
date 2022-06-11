package keyboard

import tele "gopkg.in/telebot.v3"

func NevesomoInlineKeyboard() *tele.ReplyMarkup {
	selector := &tele.ReplyMarkup{}

	btnText := selector.Data("Текст", "text", "")
	btnAudio := selector.Data("Аудио", "audio", "")
	btnTracks := selector.Data("Партии", "tracks", "")

	selector.Inline(
		selector.Row(btnText, btnAudio, btnTracks),
	)

	return selector
}
