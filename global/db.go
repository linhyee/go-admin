package global

import (
	"gorm.io/gorm"
	gormT "gorm.io/plugin/opentracing"
)

var Db *gorm.DB

func DB() *gorm.DB {
	Db.Use(gormT.New())
	return Db
}
