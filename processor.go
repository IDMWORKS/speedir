package main

import (
	"log"
	"net"

	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	packet := ber.DecodePacket(buf)

	if len(packet.Children) == 0 {
		log.Println("Error decoding asn1-ber packet: wrong port?")
		return
	}

	parsePacket(conn, packet)
}

func parsePacket(conn net.Conn, packet *ber.Packet) {
	messageID := packet.Children[0].Value.(uint64)
	response := packet.Children[1]

	if response.ClassType == ber.ClassApplication &&
		response.TagType == ber.TypeConstructed {
		if response.Tag == ldap.ApplicationBindRequest {
			handleBindRequest(messageID, response)
		}
	}
}

func handleBindRequest(messageID uint64, response *ber.Packet) {
	version := response.Children[0].Value.(uint64)
	name := response.Children[1].Value.(string)
	auth := response.Children[2]
	pass := auth.Data.String()
	log.Println("ApplicationBindRequest:",
		"messageID:", messageID,

		"LDAP version:", version,
		"username:", name,

		"password:", pass)
}
