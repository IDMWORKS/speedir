package main

import (
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
	dc := &datacontext.DataContext{DBName: dbname, DBUser: dbuser}
	setupDb(dc)
	defer dc.CloseDb()
	proc := setupProcessor(dc)
	startServers(proc)
}

func parseFlags() {
	// register flags & pointers where values will be stored
	verbosePtr := flag.Bool("verbose", false, "verbose output")
	// parse all flags - values now stored in pointers
	flag.Parse()
	// store flags for use throughout the app
	verbose = *verbosePtr
}

func setupDb(dc *datacontext.DataContext) {
	dc.InitDb()
	dc.SeedDb()
}

func setupProcessor(dc *datacontext.DataContext) *processor.Processor {
	proc := &processor.Processor{DC: dc, Verbose: verbose}
	return proc
}

func startServers(proc *processor.Processor) {
	// start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, proc.HandleRequest)
	// start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, proc.HandleRequest)
}
