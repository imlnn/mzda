package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`

	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
}

func MustLoad(svc string) Config {
	// Loading path
	configPath := os.Getenv("CONFIG_PATH_" + svc)
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// File check
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("File in CONFIG_PATH is not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", configPath)
	}

	return cfg
}
