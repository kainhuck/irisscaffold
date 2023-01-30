package configx

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

var (
	configName = "config.toml"
	FilePath   = "./res"
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

func ParseConfig[T any](cfg T) error {
	var (
		err      error
		contents []byte
	)
	filePath := filepath.Join(FilePath, configName)
	absPath, _ := filepath.Abs(filePath)

	if contents, err = os.ReadFile(absPath); err != nil {
		return fmt.Errorf("could not load configuration file: %s", err.Error())
	}
	if err = toml.Unmarshal(contents, &cfg); err != nil {
		return fmt.Errorf("could not load configuration file: %s", err.Error())
	}

	return nil
}
