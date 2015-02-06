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
	dbmap := datacontext.InitDb()
	defer dbmap.Db.Close()
	datacontext.SeedDb(dbmap)

	processor.DbMap = dbmap

	//start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, processor.HandleRequest)

	//start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, processor.HandleRequest)
}
