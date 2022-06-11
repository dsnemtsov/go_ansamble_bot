package main

import (
	"go_ensemble_bot/internal/config"
	"go_ensemble_bot/internal/process"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const configPath = "config.yml"

func main() {
	//cfg := Config(configPath)

	//pref := tele.Settings{
	//	Token:  os.Getenv("tg_ensemble_bot_token"),
	//	Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	//}

	webhook := &tele.Webhook{
		Listen:   os.Getenv("PORT"),
		Endpoint: &tele.WebhookEndpoint{PublicURL: "https://goensemblebot.herokuapp.com"},
	}

	pref := tele.Settings{
		Token:  os.Getenv("tg_ensemble_bot_token"),
		Poller: webhook,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	//go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	process.Do(b)

	b.Start()
}

func Config(path string) config.Config {
	log.Print("config initializing")
	cfg := config.Config{}
	bb, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(bb, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("config loaded")

	return cfg
}
