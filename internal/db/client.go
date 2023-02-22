package db

import "github.com/kainhuck/irisscaffold/internal/model"

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
