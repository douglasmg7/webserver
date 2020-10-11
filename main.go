package main

import (
	"webserver/db"

	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

const VERSION string = "0.1.0"
const PORT = "8080"
const NAME = "webserver"

var development bool
var production bool
var test bool

// todo - remove
var dbZunka *sql.DB

// Path for log, etc...
var workPath string

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
	log.Printf("Running in %v mode (version %s)\n", mode, VERSION)

	// Work path.
	if workPath = os.Getenv("WEBSERVER_DATA"); workPath == "" {
		panic("WEBSERVER_DATA env not defined")
	}
	os.MkdirAll(workPath, os.ModePerm)

	// Init log
	initLog()
}

func main() {
	db.Connect()
	defer db.Close()

	// Router
	router := httprouter.New()
	configRouter(router)

	// Why log.Fall work here?
	// log.Fatal(http.ListenAndServe(":"+port, router))
	log.Println("listen on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, newLogger(router)))
}
