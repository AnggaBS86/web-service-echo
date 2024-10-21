package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"web-service-echo/config"

	"github.com/jinzhu/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

// Init initializes the database connection.
func Init() {
	once.Do(func() {
		// Open the connection and assign to the global db variable
		db, err = gorm.Open(os.Getenv("DB_FACTORY"), config.GetDialect())
		if err != nil {
			log.Panic("Error connecting to the database:", err)
		}

		if sqlDB := db.DB(); err == nil {
			setupConnectionPool(sqlDB)
		} else {
			log.Panic("Failed to get database handle:", err)
		}
	})
}

// setupConnectionPool configures the connection pool settings for the database.
func setupConnectionPool(sqlDB *sql.DB) {
	maxIdleCon, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CON"))
	if err != nil {
		maxIdleCon = 100
		log.Println("Using default max idle connections:", err)
	}
	maxOpenCon, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CON"))
	if err != nil {
		maxOpenCon = 50
		log.Println("Using default max open connections:", err)
	}

	sqlDB.SetMaxIdleConns(maxIdleCon)
	sqlDB.SetMaxOpenConns(maxOpenCon)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// GetDB provides access to the db instance.
func GetDB() *gorm.DB {
	return db
}
