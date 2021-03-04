package db

import (
	"github.com/HackU-2020-vol4/back-end/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dbDriver := "mysql"
	dbUser := "hacku"
	dbPass := "password"
	dbName := "hackudb"
	dbOption := "?parseTime=true"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
	if err != nil {
		panic(err)
	}
	autoMigration(db)
}

func GetDB() (db *gorm.DB) {
	dbDriver := "mysql"
	dbUser := "hacku"
	dbPass := "password"
	dbName := "hackudb"
	dbOption := "?parseTime=true"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
	if err != nil {
		panic(err)
	}
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration(db *gorm.DB) {
	db.AutoMigrate(&entity.Publisher{})
	db.AutoMigrate(&entity.Keyword{})
	db.AutoMigrate(&entity.KeywordAssociation{})
	db.AutoMigrate(&entity.Solution{})
}
