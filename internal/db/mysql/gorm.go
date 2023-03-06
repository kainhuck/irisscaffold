package mysql

import (
	"github.com/kainhuck/irisscaffold/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Client struct {
	pool *gorm.DB
}

func NewClient(dsn string) (*Client, error) {
	pool, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	database, err := pool.DB()
	if err != nil {
		return nil, err
	}

	database.SetConnMaxLifetime(10 * time.Second)
	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(200)

	return &Client{
		pool: pool,
	}, nil
}

func (c *Client) AutoMigrate() error {
	if err := c.pool.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	return nil
}
