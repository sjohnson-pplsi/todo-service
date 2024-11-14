package config

import (
	"os"

	"github.com/cardboardrobots/config"
)

type Config struct {
	Port       int    `yml:"port" config:"PORT"`
	MongoDbUri string `yaml:"mongoDbUri" config:"MONGO_DB_URI"`
}

func LoadConfig() (Config, error) {
	return config.ReadConfigFile[Config](os.DirFS("."), "config/config.yml")
}
