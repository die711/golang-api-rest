package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rest/config"
)

var db *gorm.DB

func CreateConnection() {
	connection, _ := gorm.Open(mysql.Open(config.UrlDatabase()), &gorm.Config{})

	db = connection
}

func CloseConnection() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func CreateTables() {
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
}
