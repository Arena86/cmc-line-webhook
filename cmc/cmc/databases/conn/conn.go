package conn

import (
	mysql "cmc/cmc/databases/drivers/mysql"

	"gorm.io/gorm"
)

// Connect func
func Connect() *gorm.DB {
	return mysql.Connect()
}
