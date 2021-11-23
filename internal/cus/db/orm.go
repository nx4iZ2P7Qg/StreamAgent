package db

import (
	"StreamAgent/internal/cus/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var _db *gorm.DB

func init() {
	_db = getPostgreSQLCon()
	_db.LogMode(true)
	_db.DB().SetMaxOpenConns(100)
	_db.DB().SetMaxIdleConns(20)
}

func getPostgreSQLCon() *gorm.DB {
	user := config.C.Postgres.Username
	pass := config.C.Postgres.Password
	host := config.C.Postgres.Host
	port := config.C.Postgres.Port
	dbName := config.C.Postgres.Db
	connString := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbName)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		panic("connect to database error, " + err.Error())
	}

	schema := config.C.Postgres.Schema
	gorm.DefaultTableNameHandler = func(db *gorm.DB, table string) string {
		return schema + "." + table
	}
	return db
}

func GetDBC() *gorm.DB {
	return _db
}
