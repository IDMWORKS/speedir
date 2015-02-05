package main

const (
	listenTCPPort = 3333
	listenTLSPort = 3334
)

func main() {
	//start first TCP server in a goroutine
	go serveTCP(listenTCPPort, false, handleRequest)
	//start second TCP (TLS) server in the main thread
	serveTCP(listenTLSPort, true, handleRequest)
}
