package db

import (
	"go-template/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", "mvc.db")
	if err != nil {
		panic(err)
	}
	autoMigration()
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

func autoMigration() {
	db.AutoMigrate(&entity.Publisher{})
	db.AutoMigrate(&entity.Keyword{})
	db.AutoMigrate(&entity.KeywordAssociation{})
	db.AutoMigrate(&entity.Solution{})
}
