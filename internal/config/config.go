package config

type Config struct {
	Telegram struct {
		Host string `yaml:"host"`
	} `yaml:"telegram"`
}
