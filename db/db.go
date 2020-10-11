package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"webserver/models"

	_ "github.com/lib/pq" // postgres drive
)

var db *sqlx.DB

func Connect() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	var err error
	db, err = sqlx.Connect("postgres", "user=ws dbname=ws sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {
	defer db.Close()
}

func GetProductById(id string) (product models.Product, err error) {
	// product := models.Product{}
	err = db.Get(&product, "select * from products where id=$1", "1")
	if err != nil {
		log.Printf("[error] %v\n", err)
		return
	}
	return
}

func GetAllProduct() (products []models.Product, err error) {
	err = db.Select(&products, "select * from products where is_deleted=$1", false)
	if err != nil {
		log.Printf("[error] %v\n", err)
		return
	}
	return
}
