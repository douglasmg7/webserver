package models

import (
	"github.com/jmoiron/sqlx"
)

var pgDB *sqlx.DB

func SetDb(db *sqlx.DB) {
	pgDB = db
}
