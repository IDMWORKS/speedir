package server

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"

	"github.com/nwoolls/speedir/errors"
)

const (
	listenType = "tcp"
)

type requestHandler func(conn net.Conn)

//ServeTCP starts a TCP server on port, optionally secure with a requestHandler
func ServeTCP(port int, secure bool, handler requestHandler) {
	service := "0.0.0.0:" + strconv.Itoa(port)
	tlsFlag := "TCP"
	var err error
	var l net.Listener

	if secure {
		//cert generation tool: http://golang.org/src/crypto/tls/generate_cert.go
		cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
		errors.CheckErr(err, "Load key pair failed")

		config := tls.Config{Certificates: []tls.Certificate{cert}}
		l, err = tls.Listen(listenType, service, &config)
		tlsFlag = "TLS"
	} else {
		l, err = net.Listen(listenType, service)
	}
	errors.CheckErr(err, "TCP listen failed")

	defer l.Close()
	log.Println("Listening on", service, "("+tlsFlag+")")

	for {
		conn, err := l.Accept()
		errors.CheckErr(err, "Accept connection failed")

		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handler(conn)
	}
}
