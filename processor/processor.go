package processor

import (
	"log"
	"net"

	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
	"github.com/nwoolls/speedir/errors"
	"github.com/nwoolls/speedir/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v1"
)

//DbMap provides access to the data layer
var DbMap *gorp.DbMap

//HandleRequest handles incoming LDAPv3 requests
func HandleRequest(conn net.Conn) {
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
		switch response.Tag {
		case ldap.ApplicationBindRequest:
			handleBindRequest(messageID, response)
		default:
			log.Println("LDAPv3 app code not implemented:", response.Tag)
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
		err = bcrypt.CompareHashAndPassword([]byte(users[0].PasswordHash), []byte(password))
		if err == nil {
			log.Println("Password for user valid:", username)
		} else {
			log.Println("Password for user invalid:", username)
		}
	} else {
		log.Println("User not found:", username)
	}
}
