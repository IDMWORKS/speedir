package main

import (
	"github.com/nwoolls/speedir/datacontext"
	"github.com/nwoolls/speedir/processor"
	"github.com/nwoolls/speedir/server"
)

const (
	listenTCPPort = 3333
	listenTLSPort = 3334
)

func main() {
	//initialize DB schema
	dbmap := datacontext.InitDb()
	defer dbmap.Db.Close()
	datacontext.SeedDb(dbmap)

	//start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, processor.HandleRequest)

	//start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, processor.HandleRequest)
}
