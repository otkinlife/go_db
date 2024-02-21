package driver

import (
	"errors"
	"gorm.io/gorm"
)

const (
	DBPostGreSql = "postgres"
	DBMySql      = "mysql"
)

type GetDialectorFunc func(connectConfig *ConnectConfig) gorm.Dialector

type DB struct {
	ConnectConfig *ConnectConfig   // 连接配置
	GormConfig    *gorm.Config     // gorm配置
	PoolConfig    *PoolConfig      // 连接池配置
	dialectFunc   GetDialectorFunc // 获取dialector的方法
	pool          *Pool            // 连接池
}

func NewDB(connectConfig *ConnectConfig, gormConfig *gorm.Config, poolConfig *PoolConfig) *DB {
	return &DB{
		ConnectConfig: connectConfig,
		GormConfig:    gormConfig,
		PoolConfig:    poolConfig,
	}
}

func (d *DB) Pool() error {
	var err error
	if d.dialectFunc == nil {
		return errors.New("dialectFunc is nil")
	}
	d.registerDialect(d.dialectFunc(d.ConnectConfig))
	d.pool, err = d.NewPool()
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) RegisterDialectFunc(f GetDialectorFunc) *DB {
	d.dialectFunc = f
	return d
}
