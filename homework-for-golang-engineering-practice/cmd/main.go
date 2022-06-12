package main

import (
	"database/sql"
	"log"
	"runtime"
)

type Con struct {
	db *sql.DB
}

func NewApp(db *sql.DB) *Con {
	return &Con{
		db: db,
	}
}

func main() {
	con, cleanup, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	var ver string
	row := con.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&ver); err != nil {
		log.Fatal(err)
	}
	log.Println(ver)
	runtime.NumGoroutine()
}
