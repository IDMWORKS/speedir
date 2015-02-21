package processor

import (
	"bufio"
	"log"
	"net"

	"github.com/idmworks/speedir/datacontext"
	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
)

// DC provides access to the data layer
var DC *datacontext.DataContext

// Verbose controls the verbosity of logging
var Verbose = false

// HandleRequest handles incoming LDAPv3 requests
func HandleRequest(conn net.Conn) {
	// continuously read from the connection
	for {
		packet, err := ber.ReadPacket(bufio.NewReader(conn))

		if Verbose {
			ber.PrintPacket(packet)
		}

		if err != nil {
			defer conn.Close()
			log.Println("Error reading:", err.Error())
			return
		}

		// required to catch issues like TLS/TCP port mis-matches
		if len(packet.Children) == 0 {
			defer conn.Close()
			log.Println("Error decoding asn1-ber packet: wrong port?")
			return
		}

		parsePacket(conn, packet)
	}
}

func parsePacket(conn net.Conn, packet *ber.Packet) {
	messageID := packet.Children[0].Value.(uint64)
	request := packet.Children[1]

	if request.ClassType == ber.ClassApplication &&
		request.TagType == ber.TypeConstructed {

		switch request.Tag {
		case ldap.ApplicationBindRequest:
			handleBindRequest(conn, messageID, request)
		default:
			log.Println("LDAPv3 app code not implemented:", request.Tag)
		}

	}
}

func sendLdapResponse(conn net.Conn, packet *ber.Packet) {
	buf := packet.Bytes()

	if Verbose {
		ber.PrintPacket(packet)
	}

	for len(buf) > 0 {
		n, err := conn.Write(buf)
		if err != nil {
			log.Printf("Error Sending Message: %s\n", err)
			return
		}
		if n == len(buf) {
			break
		}
		buf = buf[n:]
	}
}

func getLdapResponse(messageID uint64, ldapResult int) *ber.Packet {
	packet := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	packet.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))
	return packet
}
