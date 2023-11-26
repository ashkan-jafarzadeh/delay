package mysql

import (
	"fmt"
	"github.com/ashkan-jafarzadeh/delay/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient(cfg config.Mysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
