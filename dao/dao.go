package dao

import (
	"github.com/jinzhu/gorm"
)

type DefaultDao struct {
	c       *dbConfig
	mySqlDB *gorm.DB
}

type dbConfig struct {
	dbType string
	dbInfo string
}
