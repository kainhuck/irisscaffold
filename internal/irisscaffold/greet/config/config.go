package config

import (
	"flag"
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"os"
)

var Cfg = new(Config)

type Config struct {
	configx.Config
}

func init() {
	// 读取配置文件
	flag.StringVar(&configx.FilePath, "c", configx.FilePath, "./sync -c configFile")
	flag.Parse()

	if err := configx.ParseConfig(Cfg); err != nil {
		fmt.Printf("read cfg file error: %s\n", err)
		os.Exit(-1)
	}
}
