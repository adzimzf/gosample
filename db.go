package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "10.164.4.106"
	port     = 5432
	user     = "af180108"
	password = "tokopedia789"
	dbname   = "tokopedia-user"
)

func getDB() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
