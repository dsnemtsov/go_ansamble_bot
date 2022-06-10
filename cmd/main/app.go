package main

import (
	"fmt"
	"go_ensemble_bot/clients/telegram"
	"go_ensemble_bot/internal/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const configPath = "config.yml"

func main() {
	token := os.Getenv("tg_ensemble_bot_token")
	cfg := Config(configPath)

	tgClient := telegram.New(cfg.Telegram.Host, token)

	fmt.Println(tgClient)

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)

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
