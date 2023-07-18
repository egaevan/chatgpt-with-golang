package config

import "os"

type Config struct {
	OpenAIApiKey string
	OpenAIUrl    string
}

func GetConfig() *Config {
	return &Config{
		OpenAIApiKey: os.Getenv("OPENAI_API_KEY"),
		OpenAIUrl:    os.Getenv("OPENAI_URL"),
	}
}
