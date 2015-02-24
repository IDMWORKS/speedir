package processor

import (
	"log"

	"github.com/mavricknz/asn1-ber"
	"github.com/mavricknz/ldap"
)

func init() {
	requestProcessors = append(requestProcessors,
		requestProcessor{
			ldapCode: ldap.ApplicationBindRequest,
			handler:  handleBindRequest,
		})
}

func handleBindRequest(proc *Processor, messageID uint64, request *ber.Packet) {
	response, result := proc.getBindResponse(messageID, request)

	if result != ldap.LDAPResultSuccess {
		defer proc.conn.Close()
	}

	proc.sendLdapResponse(response)
}

func (proc *Processor) getBindResponse(messageID uint64, request *ber.Packet) (response *ber.Packet, result int) {
	username := request.Children[1].ValueString()
	auth := request.Children[2]
	password := auth.Data.String()

	users := proc.DC.SelectUsersByUsername(username)

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
	ldapResponse := proc.buildLdapResponse(messageID, ldapResult)
	bindResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationBindResponse, nil, "Bind Response")
	bindResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagEnumerated, uint64(ldapResult), "LDAP Result"))

	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Matched DN"))
	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Error Message"))

	ldapResponse.AppendChild(bindResponse)

	return ldapResponse
}
