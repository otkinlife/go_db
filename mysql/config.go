package mysql

import (
	"github.com/otkinlife/go_db/driver"
	"gorm.io/gorm"
)

func GetGMGlobalConfig() (*driver.ConnectConfig, *gorm.Config, *driver.PoolConfig) {
	connectConfig := &driver.ConnectConfig{
		Host:     "",
		Port:     3306,
		DbName:   "bills",
		User:     "",
		Password: "",
		Charset:  "utf8mb4",
	}
	poolConfig := &driver.PoolConfig{
		MaxIdle:     300,
		MaxOpen:     500,
		MaxIdleTime: 2,
		MaxLifeTime: 30,
	}
	return connectConfig, nil, poolConfig
}

func GetGMConfig() (*driver.ConnectConfig, *gorm.Config, *driver.PoolConfig) {
	connectConfig := &driver.ConnectConfig{
		Host:     "",
		Port:     3306,
		DbName:   "",
		User:     "",
		Password: "",
		Charset:  "utf8mb4",
	}
	poolConfig := &driver.PoolConfig{
		MaxIdle:     300,
		MaxOpen:     500,
		MaxIdleTime: 2,
		MaxLifeTime: 30,
	}
	return connectConfig, nil, poolConfig
}
