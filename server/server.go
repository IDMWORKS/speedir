package server

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"

	"fmt"
)

const (
	listenType = "tcp"
)

type requestHandler func(conn net.Conn, errChan chan error)

type Server struct {
	Port    int
	Secure  bool
	Handler requestHandler
	ErrChan chan error
}

// ServeTCP starts a TCP server on port, optionally secure with a requestHandler
func (server *Server) ServeTCP() {
	listener, err := server.startListening()
	if err != nil {
		server.ErrChan <- err
		return
	}
	defer listener.Close()
	server.handleConnections(listener)
}

func (server *Server) startListening() (listener net.Listener, err error) {
	service, tlsFlag := "0.0.0.0:"+strconv.Itoa(server.Port), "TCP"

	if server.Secure {
		var config *tls.Config
		config, err = createTLSConfig()
		if err != nil {
			return nil, err
		}
		listener, err = tls.Listen(listenType, service, config)
		tlsFlag = "TLS"
	} else {
		listener, err = net.Listen(listenType, service)
	}

	if err == nil {
		log.Println("Listening on", service, "("+tlsFlag+")")
	}
	return listener, err
}

func createTLSConfig() (config *tls.Config, err error) {
	// cert generation tool: http://golang.org/src/crypto/tls/generate_cert.go
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		return nil, fmt.Errorf("Load key pair failed: %v", err)
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

func (server *Server) handleConnections(listener net.Listener) {
	// continuously accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			server.ErrChan <- fmt.Errorf("Accept connection failed: %v", err)
			continue
		}

		log.Printf("Received message %s -> %s \n",
			conn.RemoteAddr(),
			conn.LocalAddr())

		go server.Handler(conn, server.ErrChan)
	}
}
