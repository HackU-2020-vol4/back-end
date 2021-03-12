package db

import (
	"os"

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
	dbPass := os.Getenv("HACKU_PASSWORD") // 環境変数から取得　本番用
	// dbPass := "password" // ローカル用
	dbName := "hackudb"
	dbOption := "?parseTime=true"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
	if err != nil {
		panic(err)
	}
	autoMigration(db)
	defer db.Close()
}

func GetDB() (db *gorm.DB) {
	dbDriver := "mysql"
	dbUser := "hacku"
	dbPass := os.Getenv("HACKU_PASSWORD") // 環境変数から取得　本番用
	// dbPass := "password" // ローカル用
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
	db.AutoMigrate(&entity.Vote{})
}
