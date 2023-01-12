package config

import (
	bk "CleanArch/features/book/data"
	usr "CleanArch/features/user/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dc DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.DBUser, dc.DBPass, dc.DBHost, dc.DBPort, dc.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}
	return db
}

// Call this function to Creating new table in database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(usr.User{})
	db.AutoMigrate(bk.Books{})
}
