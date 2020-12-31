package dao

import (
	"gin-mvc/db"
	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {
	return db.Db
}
