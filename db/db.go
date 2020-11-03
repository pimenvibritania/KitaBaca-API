package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pimenvibritania/rest-api/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := "postgres://" + conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@" + conf.DB_HOST + ":" + conf.DB_PORT + "/" + conf.DB_NAME + "?sslmode=disable"
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic("connectionString error...")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN error...")
	}

}

func CreateCon() *sql.DB {
	return db
}
