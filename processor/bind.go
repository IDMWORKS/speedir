package processor

import (
	"log"
	"net"

	"github.com/idmworks/speedir/datacontext"
	"github.com/idmworks/speedir/errors"

	"github.com/mmitton/asn1-ber"
	"github.com/mmitton/ldap"
)

func (proc *Processor) handleBindRequest(conn net.Conn, messageID uint64, request *ber.Packet) {
	response, result := proc.getBindResponse(messageID, request)

	if result != ldap.LDAPResultSuccess {
		defer conn.Close()
	}

	proc.sendLdapResponse(conn, response)
}

func (proc *Processor) getBindResponse(messageID uint64, request *ber.Packet) (response *ber.Packet, result int) {
	username := request.Children[1].Value.(string)
	auth := request.Children[2]
	password := auth.Data.String()

	users := make(datacontext.DBUsers, 0)

	// need to patch the leaky abstraction of SQL here
	rows, err := proc.DC.DB.Query(datacontext.SqlSelectUserByUsername, username)
	errors.CheckErr(err, "Select failed")
	users.Scan(rows)

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

	response = proc.buildBindResponse(messageID, result)

	return
}

func (proc *Processor) buildBindResponse(messageID uint64, ldapResult int) *ber.Packet {
	ldapResponse := proc.getLdapResponse(messageID, ldapResult)
	bindResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationBindResponse, nil, "Bind Response")
	bindResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagEnumerated, uint64(ldapResult), "LDAP Result"))

	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Matched DN"))
	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Error Message"))

	ldapResponse.AppendChild(bindResponse)

	return ldapResponse
}
