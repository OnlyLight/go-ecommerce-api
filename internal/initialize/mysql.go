package initialize

import (
	"fmt"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.MySQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrPanic(err, "Failed to connect to MySQL database")

	global.Logger.Info("Connected to MySQL database")
	global.MDB = db

	// Set pool => open amount of connection available ready to use
	SetPool()

	// Migrate tables
	migrateTables()
}

func SetPool() {
	sqlDb, err := global.MDB.DB()
	if err != nil {
		global.Logger.Error("Failed to set pool", zap.Error(err))
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(global.Config.MySQL.MaxIdleConns))
	sqlDb.SetMaxOpenConns(global.Config.MySQL.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Config.MySQL.ConnMaxLifetime))
}

func migrateTables() {
	err := global.MDB.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}
