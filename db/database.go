package db

import (
	"log"
	"os"
	"strconv"
	"time"
	"web-service-echo/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(os.Getenv("DB_FACTORY"), config.GetDialect())
	if err != nil {
		panic("Error connection : " + err.Error())
	}

	setupConnectionPool(db)
}

func setupConnectionPool(db *gorm.DB) {
	maxIddleCon, err := strconv.Atoi(os.Getenv("DB_MAX_IDDLE_CON"))
	if err != nil {
		maxIddleCon = 100
		log.Println(err.Error())
	}
	maxOpenCon, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CON"))
	if err != nil {
		maxOpenCon = 50
		log.Println(err.Error())
	}

	db.DB().SetMaxIdleConns(maxIddleCon)
	db.DB().SetMaxOpenConns(maxOpenCon)
	db.DB().SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}
