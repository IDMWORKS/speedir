package processor

import (
	"fmt"
	"sort"
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

func (proc *Processor) processSearchRequest(messageID uint64, request *ber.Packet) (ldapResult int) {
	searchReq := &ldap.SearchRequest{
		BaseDN:       request.Children[0].ValueString(),
		Scope:        int(request.Children[1].Value.(uint64)),
		DerefAliases: int(request.Children[2].Value.(uint64)),
		SizeLimit:    int(request.Children[3].Value.(uint64)),
		TimeLimit:    int(request.Children[4].Value.(uint64)),
		TypesOnly:    request.Children[5].Value.(bool),
		Attributes:   []string{},
	}
	searchReq.Filter, _ = ldap.DecompileFilter(request.Children[6])

	subschema := false
	for _, attr := range request.Children[7].Children {
		attrName := attr.ValueString()
		if attrName == "1.1" {
			// http://www.alvestrand.no/objectid/1.1.html
			searchReq.Attributes = nil
			break
		} else if strings.EqualFold(attrName, "subschemaSubentry") {
			subschema = true
			break
		} else {
			searchReq.Attributes = append(searchReq.Attributes, attrName)
		}
	}

	sort.Strings(searchReq.Attributes)

	var ldapResponse *ber.Packet
	switch {
	case subschema:
		ldapResponse = proc.buildSubschemaResponse(messageID, *searchReq)
	case strings.EqualFold(searchReq.BaseDN, cnSchema):
		ldapResponse = proc.buildSchemaResponse(messageID, *searchReq)
	default:
		ldapResponse = proc.buildSearchEntryResponse(messageID, *searchReq)
	}

	if ldapResponse == nil {
		return ldap.LDAPResultNoSuchObject
	}
	proc.sendLdapResponse(ldapResponse)
	return ldap.LDAPResultSuccess
}

	log.Println(criteria)
func (proc *Processor) buildSearchEntryResponse(messageID uint64, searchReq ldap.SearchRequest) *ber.Packet {
	return nil
}

	log.Println(criteria)
	return nil
func (proc *Processor) buildSchemaResponse(messageID uint64, searchReq ldap.SearchRequest) *ber.Packet {
}

func (proc *Processor) buildSubschemaResponse(messageID uint64, searchReq ldap.SearchRequest) *ber.Packet {
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
