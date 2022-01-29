package models

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "orders_by"
)

func init() {
	// //set db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo) //membuka koneksi ke db
	if err != nil {
		panic(err)
	}
	DB = db
	// defer db.Close()

	err = db.Ping() //sql.Open return struct yg memiliki method Ping, maka jika sql.Open gagal db.Ping tidak bisa terksekusi
	if err != nil {
		panic(err)
	}

	fmt.Println("Sucessfully connected to database")
}
