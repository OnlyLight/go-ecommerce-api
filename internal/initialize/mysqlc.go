package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func checkErrPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQLC() {
	m := global.Config.MySQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DBName)
	// fmt.Println("connection string::", s)
	db, err := sql.Open("mysql", s)

	checkErrPanic(err, "Failed to connect to MySQL database")

	global.Logger.Info("Connected to MySQL database")
	global.Mdbc = db

	// Set pool => open amount of connection available ready to use
	SetPoolC()
}

func SetPoolC() {
	sqlDb := global.Mdbc

	sqlDb.SetMaxIdleConns(global.Config.MySQL.MaxIdleConns)
	sqlDb.SetMaxOpenConns(global.Config.MySQL.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Config.MySQL.ConnMaxLifetime))
}
