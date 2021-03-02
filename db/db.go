package db

import (
	"back-end/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("sqlite3", "mvc.db")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Room{})
	db.AutoMigrate(&entity.Publisher{})
}
