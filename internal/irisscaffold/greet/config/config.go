package config

import (
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"os"
)

type (
	CustomConfig struct {
		Name string
		Age  int
	}

	Config struct {
		configx.Config
		Custom CustomConfig
	}
)

func NewConfig() *Config {
	var cfg Config

	if err := configx.ParseConfig(&cfg); err != nil {
		fmt.Printf("read cfg file error: %v\n", err)
		os.Exit(-1)
	}

	return &cfg
}
