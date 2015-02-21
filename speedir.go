package main

import (
	"database/sql"
	"flag"

	"github.com/idmworks/speedir/datacontext"
	"github.com/idmworks/speedir/processor"
	"github.com/idmworks/speedir/server"
)

const (
	listenTCPPort = 3333
	listenTLSPort = 3334
	dbname        = "speedir"
	dbuser        = "speedir"
)

var (
	verbose = false
)

func main() {
	parseFlags()
	db := setupDb()
	defer closeDb(db)
	setupProcessor(db)
	startServers()
}

func parseFlags() {
	// register flags & pointers where values will be stored
	verbosePtr := flag.Bool("verbose", false, "verbose output")
	// parse all flags - values now stored in pointers
	flag.Parse()
	// store flags for use throughout the app
	verbose = *verbosePtr
}

func setupDb() *sql.DB {
	db := datacontext.InitDb(dbname, dbuser)
	datacontext.SeedDb(db)
	return db
}

func closeDb(db *sql.DB) {
	db.Close()
}

func setupProcessor(db *sql.DB) {
	processor.Db = db
	processor.Verbose = verbose
}

func startServers() {
	// start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, processor.HandleRequest)
	// start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, processor.HandleRequest)
}
