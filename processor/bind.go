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

func handleBindRequest(proc *Processor, messageID uint64, request *ber.Packet) error {
	response, result, err := proc.getBindResponse(messageID, request)
	if err != nil {
		return err
	}

	if result != ldap.LDAPResultSuccess {
		defer proc.conn.Close()
	}
	proc.sendLdapResponse(response)

	return nil
}

func (proc *Processor) getBindResponse(messageID uint64, request *ber.Packet) (response *ber.Packet, result int, err error) {
	username := request.Children[1].ValueString()
	auth := request.Children[2]
	password := auth.Data.String()
	result = ldap.LDAPResultProtocolError

	users, err := proc.DC.SelectUsersByUsername(username)
	if err != nil {
		return nil, result, err
	}

	if len(users) == 1 {
		log.Println("User found:", username)

		match, err := users[0].ComparePassword(password)
		switch {
		case err != nil:
			return nil, result, err
		case match:
			log.Println("Password for user valid:", username)
			result = ldap.LDAPResultSuccess
		default:
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
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))

	bindResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationBindResponse, nil, "Bind Response")
	bindResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagEnumerated, uint64(ldapResult), "LDAP Result"))

	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Matched DN"))
	bindResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Error Message"))

	ldapResponse.AppendChild(bindResponse)

	return ldapResponse
}
