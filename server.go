package main

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"
)

const (
	listenType = "tcp"
)

type requestHandler func(conn net.Conn)

func serveTCP(port int, secure bool, handler requestHandler) {
	service := "0.0.0.0:" + strconv.Itoa(port)
	tlsFlag := "TCP"
	var err error
	var l net.Listener

	if secure {
		//cert generation tool: http://golang.org/src/crypto/tls/generate_cert.go
		cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
		if err != nil {
			log.Fatal(err)
		}

		config := tls.Config{Certificates: []tls.Certificate{cert}}
		l, err = tls.Listen(listenType, service, &config)
		tlsFlag = "TLS"
	} else {
		l, err = net.Listen(listenType, service)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	log.Println("Listening on", service, "("+tlsFlag+")")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handler(conn)
	}
}
