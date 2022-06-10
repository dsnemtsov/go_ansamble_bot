package telegram

import "go_ensemble_bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
