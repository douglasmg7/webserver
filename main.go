package main

import (
	"webserver/database"
	"webserver/models"

	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq" // postgres drive
)

const VERSION string = "0.1.0"
const PORT = "8080"
const NAME = "webserver"

var development bool
var production bool
var test bool

// todo - remove
var dbZunka *sql.DB

// Path for log
var workPath string

var pgDB *sqlx.DB

// Sessions from each user.
var sessions = Sessions{
	mapUserID:      map[string]int{},
	mapSessionData: map[int]*SessionData{},
}

func init() {
	// Run mode
	var mode string
	if strings.HasPrefix(strings.ToLower(os.Getenv("RUN_MODE")), "prod") {
		production = true
		mode = "production"
	} else {
		development = true
		mode = "development"
	}

	// Work path.
	if workPath = os.Getenv("WEBSERVER_DATA"); workPath == "" {
		panic("WEBSERVER_DATA env not defined")
	}
	os.MkdirAll(workPath, os.ModePerm)

	// Init log
	initLog()

	log.Printf("Running in %v mode (version %s)\n", mode, VERSION)
}

func main() {
	pgDB = database.ConnectPostgres()
	defer database.ClosePostgres()
	models.SetDb(pgDB)

	// Router
	router := httprouter.New()
	configRouter(router)

	// Why log.Fall work here?
	// log.Fatal(http.ListenAndServe(":"+port, router))
	log.Println("Listen port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, newLogger(router)))
}
