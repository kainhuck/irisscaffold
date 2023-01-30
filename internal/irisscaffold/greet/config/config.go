package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var (
	configName = "config.toml"
	filePath   = "./res"
	Cfg        = new(Config)
)

type (
	// LogConfig logger common
	LogConfig struct {
		FilePath string
		LogLevel string
	}

	// ServiceInfo 服务地址端口
	ServiceInfo struct {
		Host string
		Port int
	}

	Config struct {
		Logger  LogConfig
		Service ServiceInfo
	}
)

func parseConfig() (*Config, error) {
	var (
		err      error
		contents []byte
		dc       Config
	)
	filePath := filepath.Join(filePath, configName)
	absPath, _ := filepath.Abs(filePath)

	if contents, err = os.ReadFile(absPath); err != nil {
		return nil, errors.New(fmt.Sprintf("could not load configuration file: %s", err.Error()))
	}
	if err = toml.Unmarshal(contents, &dc); err != nil {
		return nil, errors.New(fmt.Sprintf("could not load configuration file: %s", err.Error()))
	}

	return &dc, err
}

func init() {
	// 读取配置文件
	flag.StringVar(&filePath, "c", filePath, "./sync -c configFile")
	flag.Parse()
	var err error
	if Cfg, err = parseConfig(); err != nil {
		fmt.Printf("read cfg file error: %s\n", err)
		os.Exit(-1)
	}
}
