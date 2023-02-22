package configx

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

var (
	configName = "config.toml"
	filePath   = "./res"
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

	JwtInfo struct {
		SigKey     string
		ExpireTime int // 过期时间单位秒
	}

	Mysql struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	}

	Database struct {
		Mysql Mysql
	}

	Config struct {
		Logger   LogConfig
		Service  ServiceInfo
		Jwt      JwtInfo
		Database Database
	}
)

func (m Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DBName)
}

func ParseConfig[T any](cfg T) error {
	// 读取配置文件
	flag.StringVar(&filePath, "c", filePath, "./sync -c configFile")
	flag.Parse()
	var (
		err      error
		contents []byte
	)

	absPath, _ := filepath.Abs(filepath.Join(filePath, configName))

	if contents, err = os.ReadFile(absPath); err != nil {
		return fmt.Errorf("could not load configuration file: %s", err.Error())
	}
	if err = toml.Unmarshal(contents, &cfg); err != nil {
		return fmt.Errorf("could not load configuration file: %s", err.Error())
	}

	return nil
}