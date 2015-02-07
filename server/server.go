package server

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"

	"github.com/idmworks/speedir/errors"
)

const (
	listenType = "tcp"
)

type requestHandler func(conn net.Conn)

//ServeTCP starts a TCP server on port, optionally secure with a requestHandler
func ServeTCP(port int, secure bool, handler requestHandler) {
	listener := startListening(port, secure)
	defer listener.Close()
	handleConnections(listener, handler)
}

func startListening(port int, secure bool) net.Listener {
	service, tlsFlag := "0.0.0.0:"+strconv.Itoa(port), "TCP"
	var err error
	var listener net.Listener

	if secure {
		listener, err = tls.Listen(listenType, service, createTLSConfig())
		tlsFlag = "TLS"
	} else {
		listener, err = net.Listen(listenType, service)
	}
	errors.CheckErr(err, "TCP listen failed")

	log.Println("Listening on", service, "("+tlsFlag+")")
	return listener
}

func createTLSConfig() *tls.Config {
	//cert generation tool: http://golang.org/src/crypto/tls/generate_cert.go
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	errors.CheckErr(err, "Load key pair failed")

	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

func handleConnections(listener net.Listener, handler requestHandler) {
	for {
		conn, err := listener.Accept()
		errors.CheckErr(err, "Accept connection failed")

		log.Printf("Received message %s -> %s \n",
			conn.RemoteAddr(),
			conn.LocalAddr())

		go handler(conn)
	}
}
