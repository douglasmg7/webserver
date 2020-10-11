package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	// "webserver/models"

	_ "github.com/lib/pq" // postgres drive
)

var postgresDB *sqlx.DB

func ConnectPostgres() *sqlx.DB {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	var err error
	postgresDB, err = sqlx.Connect("postgres", "user=ws dbname=ws sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return postgresDB
}

func ClosePostgres() {
	defer postgresDB.Close()
}

// func GetProductById(id string) (product models.Product, err error) {
// // product := models.Product{}
// err = db.Get(&product, "select * from products where id=$1", "1")
// if err != nil {
// log.Printf("[error] %v\n", err)
// return
// }
// return
// }

// func GetAllProduct() (products []models.Product, err error) {
// err = db.Select(&products, "select * from products where is_deleted=$1", false)
// if err != nil {
// log.Printf("[error] %v\n", err)
// return
// }
// return
// }
