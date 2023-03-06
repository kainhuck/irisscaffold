package db

import (
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"github.com/kainhuck/irisscaffold/internal/db/mysql"
	"github.com/kainhuck/irisscaffold/internal/model"
	"os"
)

type Client interface {
	// AutoMigrate 数据迁移
	AutoMigrate() error
	// GetUserByName 查找用户
	GetUserByName(name string) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(user *model.User) error
	// UpdateUser 更新用户信息
	UpdateUser(user *model.User) error
}

func NewClient(cfg configx.Mysql) Client {
	client, err := mysql.NewClient(cfg.Dsn())
	if err != nil {
		fmt.Printf("new db client failed: %v\n", err)
		os.Exit(-1)
	}
	if err := client.AutoMigrate(); err != nil {
		fmt.Printf("autoMigrate failed: %v\n", err)
		os.Exit(-1)
	}

	return client
}
