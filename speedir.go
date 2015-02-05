package main

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"

	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
)

const (
	listenTCPPort = 3333
	listenTLSPort = 3334
	listenType    = "tcp"
)

func main() {
	//start first TCP server in a goroutine
	go serveTCP(listenTCPPort, false)
	//start second TCP (TLS) server in the main thread
	serveTCP(listenTLSPort, true)
}

func serveTCP(port int, secure bool) {
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

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	packet := ber.DecodePacket(buf)

	if len(packet.Children) == 0 {
		//handles TLS requests over non-TLS and vice-versa
		log.Println("Error decoding asn1-ber packet: wrong port?")
		return
	}

	messageID := packet.Children[0].Value.(uint64)
	response := packet.Children[1]

	if response.ClassType == ber.ClassApplication &&
		response.TagType == ber.TypeConstructed {

		if response.Tag == ldap.ApplicationBindRequest {
			version := response.Children[0].Value.(uint64)
			name := response.Children[1].Value.(string)
			auth := response.Children[2]
			pass := auth.Data.String()
			log.Println("ApplicationBindRequest:",
				"messageID:", messageID,
				//"response", response,
				"LDAP version:", version,
				"username:", name,
				//"auth", auth,
				"password:", pass)
		}
	}
}
