package main

import (
	"github.com/idmworks/speedir/datacontext"
	"github.com/idmworks/speedir/processor"
	"github.com/idmworks/speedir/server"
)

const (
	listenTCPPort = 3333
	listenTLSPort = 3334
)

func main() {
	//initialize DB schema
	dbmap := datacontext.InitDb("speedir", "speedir")
	//close DB when app exits
	defer dbmap.Db.Close()
	//seed DB data
	datacontext.SeedDb(dbmap)

	//give processor access to data
	processor.DbMap = dbmap

	//start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, processor.HandleRequest)
	//start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, processor.HandleRequest)
}
