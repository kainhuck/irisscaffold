package mysql

import (
	"github.com/kainhuck/irisscaffold/internal/model"
)

func (c *Client) GetUserByName(name string) (user *model.User, err error) {
	err = c.pool.Where(&model.User{Username: name}).First(&user).Error

	return
}

func (c *Client) CreateUser(user *model.User) (err error) {
	err = c.pool.Create(user).Error

	return
}

func (c *Client) UpdateUser(user *model.User) (err error) {
	err = c.pool.Save(user).Error

	return
}
