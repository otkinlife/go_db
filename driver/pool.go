package driver

import (
	"database/sql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Pool struct {
	config *PoolConfig
	db     *gorm.DB
}

// PoolConfig 连接池配置
type PoolConfig struct {
	MaxIdle     int           //空闲连接
	MaxOpen     int           //最大连接
	MaxIdleTime time.Duration //最大空闲时间
	MaxLifeTime time.Duration //最大可复用时间

	gormConfig gorm.Config
	dialector  gorm.Dialector
}

// NewPool 初始化一个新的连接池
func (d *DB) NewPool() (*Pool, error) {
	theDb, err := gorm.Open(d.PoolConfig.dialector, &d.PoolConfig.gormConfig)
	if err != nil {
		log.Print("NewPool Error", err)
		return nil, err
	}
	sqlDB, err := theDb.DB()
	if err != nil {
		log.Print("NewPool Error", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(d.PoolConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(d.PoolConfig.MaxOpen)
	sqlDB.SetConnMaxIdleTime(d.PoolConfig.MaxIdleTime)
	sqlDB.SetConnMaxLifetime(d.PoolConfig.MaxLifeTime)
	pool := &Pool{
		db:     theDb,
		config: d.PoolConfig,
	}
	return pool, nil
}

// GetDb 获取GormDB实例
func (d *DB) GetDb() *gorm.DB {
	sqlDB, err := d.pool.db.DB()
	if err != nil {
		log.Print("GetDb Error", err)
		return nil
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Print("GetDb Error", "Ping Error", err)
		// 连接无效，重新初始化连接池
		newPool, err := d.NewPool()
		if err != nil {
			log.Print("GetDb Error", "ReInit Error", err)
		}
		d.pool.db = newPool.db
	}

	return d.pool.db
}

func (d *DB) Stats() sql.DBStats {
	sqlDB, _ := d.pool.db.DB()
	return sqlDB.Stats()
}

func (d *DB) Close() error {
	sqlDB, _ := d.pool.db.DB()
	return sqlDB.Close()
}

func (d *DB) registerDialect(dialect gorm.Dialector) {
	d.PoolConfig.dialector = dialect
}
