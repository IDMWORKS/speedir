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
	response, result := getBindResponse(messageID, request)

	if result != ldap.LDAPResultSuccess {
		defer conn.Close()
	}

	sendLdapResponse(conn, response)
}

func getBindResponse(messageID uint64, request *ber.Packet) (response *ber.Packet, result int) {
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

	result = ldap.LDAPResultProtocolError

	if len(users) == 1 {
		log.Println("User found:", username)

		if users[0].ComparePassword(password) {
			log.Println("Password for user valid:", username)
			result = ldap.LDAPResultSuccess
		} else {
			log.Println("Password for user invalid:", username)
			result = ldap.LDAPResultInvalidCredentials
		}

	} else {
		log.Println("User not found:", username)
		result = ldap.LDAPResultInvalidCredentials
	}

	response = buildBindResponse(messageID, result)

	return
}

func buildBindResponse(messageID uint64, ldapResult int) *ber.Packet {
	ldapResponse := getLdapResponse(messageID, ldapResult)
	bindResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationBindResponse, nil, "Bind Response")
	bindResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagEnumerated, uint64(ldapResult), "LDAP Result"))

	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Matched DN"))
	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Error Message"))

	ldapResponse.AppendChild(bindResponse)

	return ldapResponse
}
