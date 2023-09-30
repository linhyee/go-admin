package db

import (
	"blog/global"
	"blog/pkg/setting"
	"fmt"
	"reflect"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDB(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)

	log := logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: log,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不加s
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	db.Callback().Create().Before("gorm:create").Register("ex_plugin:update_created_at", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("ex_plugin:update_updated_at", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	if global.ServerSetting.IsTracing && global.Tracer != nil {
	}
	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		now := time.Now().Unix()
		createdAt := db.Statement.Schema.LookUpField("created_at")
		if createdAt == nil {
			return
		}
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Struct:
			createdAt.Set(db.Statement.Context, db.Statement.ReflectValue, float64(now))
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {

}

func deleteCallback(db *gorm.DB) {

}

func addExtraSpaceIfExists(str string) string {
	return ""
}
