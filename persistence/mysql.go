package persistence

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	//username of user(root)
	username = "root"
	//password of user(root)
	password = "secretpassword"
	hostname = "mysql-container:3306"
	dbname = "players"
)

func Init() {
	var err error
	DB, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Fatal(err)
	}
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
