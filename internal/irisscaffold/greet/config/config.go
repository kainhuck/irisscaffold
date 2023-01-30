package config

import (
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"os"
)

var Cfg = new(Config)

type Config struct {
	configx.Config
}

func init() {
	if err := configx.ParseConfig(Cfg); err != nil {
		fmt.Printf("read cfg file error: %s\n", err)
		os.Exit(-1)
	}
}
