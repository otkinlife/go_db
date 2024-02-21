package mysql

import (
	"github.com/otkinlife/go_db/driver"
	"testing"
)

type item struct {
	Uid int `gorm:"column:uid"`
}

func TestNewPostgreSQLPool(t *testing.T) {
	db, err := NewConnPool(
		&driver.ConnectConfig{
			Protocol: "tcp",
			Host:     "127.0.0.1",
			Port:     3306,
			DbName:   "test",
			User:     "root",
			Password: "123456",
			Charset:  "utf8",
		},
		nil,
		&driver.PoolConfig{
			MaxIdle:     10,
			MaxOpen:     10,
			MaxIdleTime: 3600,
			MaxLifeTime: 3600,
		})
	if err != nil {
		t.Error(err)
		return
	}
	count := int64(0)
	err = db.GetDb().Table("test").Count(&count).Error
	if err != nil {
		t.Error(err)
		return
	}
}
