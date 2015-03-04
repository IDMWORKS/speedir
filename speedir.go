package main

import (
	"flag"
	"log"

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
	dc, err := setupDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dc.CloseDb()
	proc := setupProcessor(dc)
	if err = startServers(proc); err != nil {
		log.Fatal(err)
	}
}

func parseFlags() {
	// register flags & pointers where values will be stored
	verbosePtr := flag.Bool("verbose", false, "verbose output")
	// parse all flags - values now stored in pointers
	flag.Parse()
	// store flags for use throughout the app
	verbose = *verbosePtr
}

func setupDb() (dc *datacontext.DataContext, err error) {
	dc = &datacontext.DataContext{DBName: dbname, DBUser: dbuser}
	if err := dc.InitDb(); err != nil {
		return nil, err
	}
	dc.SeedDb()
	return dc, nil
}

func setupProcessor(dc *datacontext.DataContext) *processor.Processor {
	proc := &processor.Processor{DC: dc, Verbose: verbose}
	return proc
}

func startServers(proc *processor.Processor) error {
	errChan := make(chan error)

	// start first TCP (TLS) server in a goroutine
	tlsServer := &server.Server{
		Port:    listenTLSPort,
		Secure:  true,
		Handler: proc.HandleRequest,
		ErrChan: errChan,
	}
	go tlsServer.ServeTCP()

	// start second TCP server in a goroutine
	tcpServer := &server.Server{
		Port:    listenTCPPort,
		Secure:  false,
		Handler: proc.HandleRequest,
		ErrChan: errChan,
	}
	go tcpServer.ServeTCP()

	// block waiting for either server to error
	for {
		select {
		case err := <-errChan:
			if err == processor.ErrDecodingASN1 {
				log.Println(err)
			} else {
				return err
			}
		}
	}
}
