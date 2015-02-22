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
	dc := setupDb()
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

func setupDb() *datacontext.DataContext {
	dc := &datacontext.DataContext{DBName: dbname, DBUser: dbuser}
	dc.InitDb()
	dc.SeedDb()
	return dc
}

func setupProcessor(dc *datacontext.DataContext) *processor.Processor {
	proc := &processor.Processor{DC: dc, Verbose: verbose}
	return proc
}

func startServers(proc *processor.Processor) {
	// start first TCP server in a goroutine
	tcpServer := &server.Server{Port: listenTCPPort, Secure: false, Handler: proc.HandleRequest}
	go tcpServer.ServeTCP()
	// start second TCP (TLS) server in the main thread
	tlsServer := &server.Server{Port: listenTLSPort, Secure: true, Handler: proc.HandleRequest}
	tlsServer.ServeTCP()
}
