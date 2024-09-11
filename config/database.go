package config

import (
	"fmt"
	"os"
)

type Database struct {
	User         string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

func (d *Database) SetUser(user string) {
	d.User = user
}

func (d *Database) GetUser() string {
	return d.User
}

func (d *Database) SetPassword(password string) {
	d.Password = password
}

func (d *Database) GetPassword() string {
	return d.Password
}

func (d *Database) SetHost(host string) {
	d.Host = host
}

func (d *Database) GetHost() string {
	return d.Host
}

func (d *Database) SetPort(port int) {
	d.Port = port
}

func (d *Database) GetPort() int {
	return d.Port
}

func (d *Database) SetDatabaseName(databaseName string) {
	d.DatabaseName = databaseName
}

func (d *Database) GetDatabaseName() string {
	return d.DatabaseName
}

func GetDialect() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
}
