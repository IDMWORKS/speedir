package processor

import (
	"bufio"
	"log"
	"net"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
	"gopkg.in/gorp.v1"
)

//DbMap provides access to the data layer
var DbMap *gorp.DbMap

//HandleRequest handles incoming LDAPv3 requests
func HandleRequest(conn net.Conn) {
	defer conn.Close()

	packet, err := ber.ReadPacket(bufio.NewReader(conn))
	if err != nil {
		log.Println("Error reading:", err.Error())
		return
	}

	parsePacket(conn, packet)
}

func parsePacket(conn net.Conn, packet *ber.Packet) {
	messageID := packet.Children[0].Value.(uint64)
	request := packet.Children[1]

	if request.ClassType == ber.ClassApplication &&
		request.TagType == ber.TypeConstructed {

		switch request.Tag {
		case ldap.ApplicationBindRequest:
			handleBindRequest(messageID, request)
		default:
			log.Println("LDAPv3 app code not implemented:", request.Tag)
			ber.PrintPacket(packet)
		}

	}
}

func handleBindRequest(messageID uint64, response *ber.Packet) {
	version := response.Children[0].Value.(uint64)
	username := response.Children[1].Value.(string)
	auth := response.Children[2]
	password := auth.Data.String()

	log.Println("\nApplicationBindRequest:",
		"\n\tmessageID:", messageID,
		"\n\tLDAP version:", version,
		"\n\tusername:", username,
		"\n\tpassword:", "********")

	var users []models.User
	_, err := DbMap.Select(&users, "select * from users where username=$1", username)
	errors.CheckErr(err, "Select failed")

	if len(users) == 1 {
		log.Println("User found:", username)

		if users[0].ComparePassword(password) {
			log.Println("Password for user valid:", username)
		} else {
			log.Println("Password for user invalid:", username)
		}

	} else {
		log.Println("User not found:", username)
	}
}
