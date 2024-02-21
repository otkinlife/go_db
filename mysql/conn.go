package mysql

import (
	"fmt"
	"github.com/otkinlife/go_db/driver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnPool(connectConfig *driver.ConnectConfig, gormConfig *gorm.Config, poolConfig *driver.PoolConfig) (*driver.DB, error) {
	newDB := driver.NewDB(connectConfig, gormConfig, poolConfig)
	err := newDB.RegisterDialectFunc(GetDialector).Pool()
	if err != nil {
		return nil, err
	}
	return newDB, err
}

func GetDialector(config *driver.ConnectConfig) gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DbName, config.Charset))
}
