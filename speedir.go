package main

import (
	"flag"

	"github.com/idmworks/speedir/datacontext"
	"github.com/idmworks/speedir/processor"
	"github.com/idmworks/speedir/server"
	"gopkg.in/gorp.v1"
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
	dbmap := setupDb()
	defer closeDb(dbmap)
	setupProcessor(dbmap)
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

func setupDb() *gorp.DbMap {
	dbmap := datacontext.InitDb(dbname, dbuser)
	datacontext.SeedDb(dbmap)
	return dbmap
}

func closeDb(dbmap *gorp.DbMap) {
	dbmap.Db.Close()
}

func setupProcessor(dbmap *gorp.DbMap) {
	processor.DbMap = dbmap
	processor.Verbose = verbose
}

func startServers() {
	// start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, processor.HandleRequest)
	// start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, processor.HandleRequest)
}
