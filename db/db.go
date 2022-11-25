package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "123"
	dbname   = "postgres"
)

type Pool struct {
	ID       int
	Login    string
	Password string
	Access   string
}

func Conectdb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
