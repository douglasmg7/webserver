package main

import (
	"fmt"
	"webserver/database"
	"webserver/models"

	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis/v7"
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

// Dbs
var pgDB *sqlx.DB
var redisDB *redis.Client

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
	// Init Postgres
	pgDB = database.ConnectPostgres()
	defer database.ClosePostgres()
	models.SetDb(pgDB)

	// Init Redis
	ConnectRedis()

	// Router
	router := httprouter.New()
	configRouter(router)

	// Categories
	models.UpdateCategories()

	log.Println("Listen port", PORT)
	Error(fmt.Errorf("Some thing wrong. %v", "Just a test"))
	debug("Teste: %v, module: %v", "a", 2)
	warn("Atenção pro horário: %v", "11:11")
	trace("Já são: %v", "11:31")

	// Why log.Fall work here?
	// log.Fatal(http.ListenAndServe(":"+port, router))
	log.Fatal(http.ListenAndServe(":"+PORT, newLogger(newSessionMiddleware(router))))
	// log.Fatal(http.ListenAndServe(":"+PORT, newLogger(router)))
}
