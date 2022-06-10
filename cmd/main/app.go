package main

import (
	"go_ensemble_bot/internal/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	log.Print("config initializing")
	cfg := config.Config{}
	bb, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(bb, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("config loaded")
}
