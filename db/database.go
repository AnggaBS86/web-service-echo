package db

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
	"web-service-echo/config"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
	mu  sync.Mutex // Mutex to prevent race conditions
)

// Init initializes the database connection with a mutex to prevent race conditions.
func Init() {
	mu.Lock() // Lock the mutex to ensure no race conditions
	defer mu.Unlock()

	// Initialize the database
	db, err = gorm.Open(os.Getenv("DB_FACTORY"), config.GetDialect())
	if err != nil {
		panic("Error connection : " + err.Error())
	}

	// Setup the connection pool
	setupConnectionPool(db)
}

// setupConnectionPool configures the connection pool settings for the database.
func setupConnectionPool(db *gorm.DB) {
	maxIdleCon, err := strconv.Atoi(os.Getenv("DB_MAX_IDDLE_CON"))
	if err != nil {
		maxIdleCon = 100
		log.Println(err.Error())
	}
	maxOpenCon, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CON"))
	if err != nil {
		maxOpenCon = 50
		log.Println(err.Error())
	}

	db.DB().SetMaxIdleConns(maxIdleCon)
	db.DB().SetMaxOpenConns(maxOpenCon)
	db.DB().SetConnMaxLifetime(time.Hour)
}

// GetDB provides thread-safe access to the db instance.
func GetDB() *gorm.DB {
	mu.Lock() // Lock the mutex to ensure thread safety
	defer mu.Unlock()

	return db
}
