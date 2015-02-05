package main

import (
	"fmt"
	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
	"net"
	"os"
)

const (
	ListenPort = "3333"
	ListenType = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(ListenType, ":"+ListenPort)
	if err != nil {
		log.Fatal(err)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on :" + ListenPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		//logs an incoming message
		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
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
			pass := ""

			//tag == 0 (I think) indicates simple auth
			//tag == 3 should be SASL once we support SSL/TLS
			if auth.Tag == 0 {
				pass = auth.Data.String()
				fmt.Println("ApplicationBindRequest:",
					"messageID:", messageID,
					//"response", response,
					"LDAP version:", version,
					"username:", name,
					//"auth", auth,
					"password:", pass)
			}
		}

	}
}
