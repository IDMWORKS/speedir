package main

import (
	"github.com/nwoolls/speedir/processor"
	"github.com/nwoolls/speedir/server"
)

const (
	listenTCPPort = 3333
	listenTLSPort = 3334
)

func main() {
	//start first TCP server in a goroutine
	go server.ServeTCP(listenTCPPort, false, processor.HandleRequest)

	//start second TCP (TLS) server in the main thread
	server.ServeTCP(listenTLSPort, true, processor.HandleRequest)
}
