package processor

import (
	"bufio"
	"log"
	"net"

	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
	"gopkg.in/gorp.v1"
)

// DbMap provides access to the data layer
var DbMap *gorp.DbMap

// HandleRequest handles incoming LDAPv3 requests
func HandleRequest(conn net.Conn) {
	// continuously read from the connection
	for {
		packet, err := ber.ReadPacket(bufio.NewReader(conn))
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
			ber.PrintPacket(packet)
		}

	}
}
