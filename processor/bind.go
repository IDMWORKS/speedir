package processor

import (
	"log"
	"net"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
)

func handleBindRequest(conn net.Conn, messageID uint64, request *ber.Packet) {
	version := request.Children[0].Value.(uint64)
	username := request.Children[1].Value.(string)
	auth := request.Children[2]
	password := auth.Data.String()

	log.Println("\nBindRequest:",
		"\n\tmessageID:", messageID,
		"\n\tLDAP version:", version,
		"\n\tusername:", username,
		"\n\tpassword:", "********")

	var users []models.User
	_, err := DbMap.Select(&users, "select * from users where username=$1", username)
	errors.CheckErr(err, "Select failed")

	bindResult := ldap.LDAPResultProtocolError

	if len(users) == 1 {
		log.Println("User found:", username)

		if users[0].ComparePassword(password) {
			log.Println("Password for user valid:", username)
			bindResult = ldap.LDAPResultSuccess
		} else {
			log.Println("Password for user invalid:", username)
			bindResult = ldap.LDAPResultInvalidCredentials
		}

	} else {
		log.Println("User not found:", username)
		bindResult = ldap.LDAPResultInvalidCredentials
	}

	if bindResult != ldap.LDAPResultSuccess {
		defer conn.Close()
	}

	sendBindResponse(conn, messageID, bindResult)
}

func sendBindResponse(conn net.Conn, messageID uint64, ldapResult int) {
	packet := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	packet.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))
	bindResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationBindResponse, nil, "Bind Response")
	bindResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagEnumerated, uint64(ldapResult), "LDAP Result"))

	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Matched DN"))
	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Error Message"))

	packet.AppendChild(bindResponse)

	buf := packet.Bytes()

	log.Println("Sending BindResponse:")
	ber.PrintPacket(packet)

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
