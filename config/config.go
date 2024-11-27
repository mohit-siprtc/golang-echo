package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type MongoConfig struct {
	MONGODB_URI   string `env:"MONGO_URI" envDefault:""`
	MONGO_DB_NAME string `env:"MONGO_DB_NAME" envDefault:""`
}

func LoadMongoConfig() *MongoConfig {
	cfg := &MongoConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	return cfg
}
