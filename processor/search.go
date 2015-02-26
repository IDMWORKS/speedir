package processor

import (
	"log"
	"strings"

	"github.com/mavricknz/asn1-ber"
	"github.com/mavricknz/ldap"
)

const (
	cnSchema = "cn=schema"
)

func init() {
	requestProcessors = append(requestProcessors,
		requestProcessor{
			ldapCode: ldap.ApplicationSearchRequest,
			handler:  handleSearchRequest,
		})
}

func handleSearchRequest(proc *Processor, messageID uint64, request *ber.Packet) {
	ldapResult := proc.processSearchRequest(messageID, request)
	ldapResponse := proc.buildSearchDoneResponse(messageID, ldapResult)
	proc.sendLdapResponse(ldapResponse)
}

type searchCriteria struct {
	baseDn       string
	scope        uint64
	derefAliases uint64
	sizeLimit    uint64
	timeLimit    uint64
	typesOnly    bool
	filter       string
	attributes   []string
}

func (proc *Processor) processSearchRequest(messageID uint64, request *ber.Packet) (ldapResult int) {
	criteria := &searchCriteria{
		baseDn:       request.Children[0].ValueString(),
		scope:        request.Children[1].Value.(uint64),
		derefAliases: request.Children[2].Value.(uint64),
		sizeLimit:    request.Children[3].Value.(uint64),
		timeLimit:    request.Children[4].Value.(uint64),
		typesOnly:    request.Children[5].Value.(bool),
	}
	criteria.filter, _ = ldap.DecompileFilter(request.Children[6])

	subschema := false
	for _, attr := range request.Children[7].Children {
		attrName := attr.ValueString()
		if attrName == "1.1" {
			// http://www.alvestrand.no/objectid/1.1.html
			criteria.attributes = nil
			break
		} else if strings.EqualFold(attrName, "subschemaSubentry") {
			subschema = true
			break
		} else {
			criteria.attributes = append(criteria.attributes, attrName)
		}
	}

	var ldapResponse *ber.Packet
	switch {
	case subschema:
		ldapResponse = proc.buildSubschemaResponse(messageID, *criteria)
	case strings.EqualFold(criteria.baseDn, cnSchema):
		ldapResponse = proc.buildSchemaResponse(messageID, *criteria)
	default:
		ldapResponse = proc.buildSearchEntryResponse(messageID, *criteria)
	}

	if ldapResponse == nil {
		return ldap.LDAPResultNoSuchObject
	}
	proc.sendLdapResponse(ldapResponse)
	return ldap.LDAPResultSuccess
}

func (proc *Processor) buildSearchEntryResponse(messageID uint64, criteria searchCriteria) *ber.Packet {
	log.Println(criteria)
	return nil
}

func (proc *Processor) buildSchemaResponse(messageID uint64, criteria searchCriteria) *ber.Packet {
	log.Println(criteria)
	return nil
}

func (proc *Processor) buildSubschemaResponse(messageID uint64, criteria searchCriteria) *ber.Packet {
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))
	searchResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationSearchResultEntry, nil, "Search Result Entry")
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "objectName	LDAPDN"))
	attributesPacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Attributes")
	attributesPacket.AppendChild(buildAttributePacket("subschemaSubentry", cnSchema))
	searchResponse.AppendChild(attributesPacket)
	ldapResponse.AppendChild(searchResponse)
	return ldapResponse
}

func buildAttributePacket(name string, values ...string) *ber.Packet {
	attributePacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Attribute")
	attributePacket.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, name, ""))
	valuesPacket := buildValuesPacket(values)
	attributePacket.AppendChild(valuesPacket)
	return attributePacket
}

func buildValuesPacket(values []string) *ber.Packet {
	valuesPacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSet, nil, "Values")
	for _, value := range values {
		valuesPacket.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, value, ""))
	}
	return valuesPacket
}

func (proc *Processor) buildSearchDoneResponse(messageID uint64, ldapResult int) *ber.Packet {
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))
	searchResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationSearchResultDone, nil, "Search Result Done")
	searchResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagEnumerated, uint64(ldapResult), "LDAP Result"))
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Matched DN"))
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "Error Message"))
	ldapResponse.AppendChild(searchResponse)
	return ldapResponse
}
