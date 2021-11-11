package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

// Connect func
func Connect() *gorm.DB {
	var hasConnection = false
	if gormDB != nil {
		mysqlDB, errMysql := gormDB.DB()
		if errMysql == nil {
			errMysql = mysqlDB.Ping()
			if errMysql == nil {
				hasConnection = true
				return gormDB
			}
		}
	}
	if !hasConnection {
		connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		for i := 1; i <= 3; i++ {
			sqlDB, err := sql.Open("mysql", connString)
			gormDB, err := gorm.Open(mysql.New(mysql.Config{
				Conn: sqlDB,
			}), &gorm.Config{})
			//}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

			// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
			sqlDB.SetMaxIdleConns(10)

			// SetMaxOpenConns sets the maximum number of open connections to the database.
			sqlDB.SetMaxOpenConns(100)

			// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
			sqlDB.SetConnMaxLifetime(time.Hour)
			if err == nil {
				return gormDB
			}
			if i == 3 {
				panic(err.Error())
			}
		}
	}
	return nil
}
