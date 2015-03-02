package processor

import (
	"fmt"
	"sort"
	"strings"

	"github.com/idmworks/speedir/datacontext"
	"github.com/idmworks/speedir/models"
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
	proc.sendSearchDoneResponse(messageID, ldapResult)
}

func (proc *Processor) sendSearchDoneResponse(messageID uint64, ldapResult int) {
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
	namingContexts := false

	for _, attr := range request.Children[7].Children {
		attrName := attr.ValueString()
		switch {
		case attrName == "1.1":
			// http://www.alvestrand.no/objectid/1.1.html
			searchReq.Attributes = nil
			break
		case strings.EqualFold(attrName, models.SubschemaSubentryAttribute):
			subschema = true
			break
		case strings.EqualFold(attrName, models.NamingContextsAttribute):
			namingContexts = true
			break
		default:
			searchReq.Attributes = append(searchReq.Attributes, attrName)
		}
	}

	sort.Strings(searchReq.Attributes)

	// TODO: derefFindingBaseObj vs derefInSearching

	ldapResult = ldap.LDAPResultNoSuchObject

	switch {
	case subschema:
		ldapResult = proc.sendSubschemaResponse(messageID, *searchReq)
	case namingContexts:
		ldapResult = proc.sendNamingContextsResponse(messageID, *searchReq)
	case strings.EqualFold(searchReq.BaseDN, cnSchema):
		ldapResult = proc.sendSchemaResponse(messageID, *searchReq)
	default:
		ldapResult = proc.sendSearchEntryResponse(messageID, *searchReq)
	}

	return ldapResult
}

func (proc *Processor) sendSearchEntryResponse(messageID uint64, searchReq ldap.SearchRequest) (ldapResult int) {
	var entries datacontext.DBEntries
	switch searchReq.Scope {
	case ldap.ScopeBaseObject:
		entries = proc.DC.SelectEntriesByDN(searchReq.BaseDN)
	case ldap.ScopeSingleLevel:
		entries = proc.DC.SelectEntriesByParent(searchReq.BaseDN)
	case ldap.ScopeWholeSubtree:
		entries = proc.DC.SelectEntryTreeByParent(searchReq.BaseDN)
	}
	if (searchReq.Scope == ldap.ScopeBaseObject) && (len(entries) == 0) {
		return ldap.LDAPResultNoSuchObject
	}

	for _, entry := range entries {
		proc.processSearchEntryResult(messageID, entry)
	}

	return ldap.LDAPResultSuccess
}

func (proc *Processor) processSearchEntryResult(messageID uint64, entry *datacontext.DBEntry) {
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))

	searchResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationSearchResultEntry, nil, "Search Result Entry")
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, entry.DN, "objectName	LDAPDN"))

	attributesPacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Attributes")

	// add attributes from entry.UserValues
	// TODO: does not handle mult-value attributes - need to group
	for key, value := range entry.UserValues {
		attributesPacket.AppendChild(buildAttributePacket(key, value))
	}

	// add attributes from entry.OperValues - may not be necessary since we normalized subschema
	// TODO: does not handle mult-value attributes
	for key, value := range entry.OperValues {
		attributesPacket.AppendChild(buildAttributePacket(key, value))
	}

	// we don't store "SUP top" in the DB - it's just a NULL SUP
	values := []string{models.TopClass}
	// add attributes from entry.Classes
	for _, objectClass := range entry.Classes {
		values = append(values, objectClass)
	}
	attributesPacket.AppendChild(buildAttributePacket(models.ObjectClassAttribute, values...))

	searchResponse.AppendChild(attributesPacket)
	ldapResponse.AppendChild(searchResponse)

	proc.sendLdapResponse(ldapResponse)
}

func (proc *Processor) sendSchemaResponse(messageID uint64, searchReq ldap.SearchRequest) (ldapResult int) {
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))

	searchResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationSearchResultEntry, nil, "Search Result Entry")
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, cnSchema, "objectName	LDAPDN"))

	attributesPacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Attributes")

	appendSchemaAttributes(attributesPacket)

	atts := searchReq.Attributes

	if i := sort.SearchStrings(atts, models.LDAPSyntaxesAttribute); i < len(atts) {
		proc.appendSyntaxAttributes(attributesPacket)
	}

	if i := sort.SearchStrings(atts, models.ObjectClassesAttribute); i < len(atts) {
		proc.appendObjectClassAttributes(attributesPacket)
	}

	if i := sort.SearchStrings(atts, models.MatchingRulesAttribute); i < len(atts) {
		proc.appendMatchingRuleAttributes(attributesPacket)
	}

	if i := sort.SearchStrings(atts, "attributesTypes"); i < len(atts) {
		proc.appendAttributeTypeAttributes(attributesPacket)
	}

	searchResponse.AppendChild(attributesPacket)
	ldapResponse.AppendChild(searchResponse)

	proc.sendLdapResponse(ldapResponse)

	return ldap.LDAPResultSuccess
}

func appendSchemaAttributes(attributesPacket *ber.Packet) {
	attributesPacket.AppendChild(buildAttributePacket(models.CommonNameAttribute, "schema"))
	attributesPacket.AppendChild(buildAttributePacket(models.ObjectClassAttribute, models.TopClass))
	attributesPacket.AppendChild(buildAttributePacket(models.ObjectClassAttribute, "ldapSubentry"))
	attributesPacket.AppendChild(buildAttributePacket(models.ObjectClassAttribute, models.SubschemaClass))
}

func (proc *Processor) appendMatchingRuleAttributes(attributesPacket *ber.Packet) {
	rules := proc.DC.SelectAllMatchingRules()
	values := []string{}
	for _, rule := range rules {
		values = append(values, fmt.Sprintf(
			"( %s NAME '%s' SYNTAX %s )",
			rule.OID,
			rule.Name,
			rule.Syntax))
	}
	attributesPacket.AppendChild(buildAttributePacket(models.MatchingRulesAttribute, values...))
}

func (proc *Processor) appendAttributeTypeAttributes(attributesPacket *ber.Packet) {
	rules := proc.DC.SelectAllAttributeTypes()
	values := []string{}
	for _, rule := range rules {
		values = append(values, fmt.Sprintf(
			"( %s NAME %s%s%s%s%s%s%s)",
			rule.OID,
			rule.NamesString(),
			rule.SuperString(),
			rule.EqualityMatchString(),
			rule.SubstrMatchString(),
			rule.OrderingMatchString(),
			rule.SyntaxString(),
			rule.FlagsString()))
	}
	attributesPacket.AppendChild(buildAttributePacket(models.AttributeTypesAttribute, values...))
}

func (proc *Processor) appendSyntaxAttributes(attributesPacket *ber.Packet) {
	syntaxes := proc.DC.SelectAllSyntaxes()
	values := []string{}
	for _, syntax := range syntaxes {
		values = append(values, fmt.Sprintf(
			"( %s DESC '%s' )",
			syntax.OID,
			syntax.Description))
	}
	attributesPacket.AppendChild(buildAttributePacket(models.LDAPSyntaxesAttribute, values...))
}

func (proc *Processor) appendObjectClassAttributes(attributesPacket *ber.Packet) {
	objectClasses := proc.DC.SelectAllObjectClasses()
	values := []string{}
	for _, objectClass := range objectClasses {
		values = append(values, fmt.Sprintf(
			"( %s NAME '%s' SUP %s %s %s %s )",
			objectClass.OID,
			objectClass.Name,
			objectClass.Super.String,
			objectClass.FlagsString(),
			objectClass.MustString(),
			objectClass.MayString()),
		)
	}
	attributesPacket.AppendChild(buildAttributePacket(models.ObjectClassesAttribute, values...))
}

func (proc *Processor) sendNamingContextsResponse(messageID uint64, searchReq ldap.SearchRequest) (ldapResult int) {
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))

	searchResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationSearchResultEntry, nil, "Search Result Entry")
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "objectName	LDAPDN"))

	attributesPacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Attributes")

	proc.appendNamingContextAttributes(attributesPacket)

	searchResponse.AppendChild(attributesPacket)
	ldapResponse.AppendChild(searchResponse)

	proc.sendLdapResponse(ldapResponse)

	return ldap.LDAPResultSuccess
}

func (proc *Processor) appendNamingContextAttributes(attributesPacket *ber.Packet) {
	entries := proc.DC.SelectAllNamingContexts()
	values := []string{}
	for _, entry := range entries {
		values = append(values, entry.DN)
	}
	attributesPacket.AppendChild(buildAttributePacket(models.NamingContextsAttribute, values...))
}

func (proc *Processor) sendSubschemaResponse(messageID uint64, searchReq ldap.SearchRequest) (ldapResult int) {
	ldapResponse := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "LDAP Response")
	ldapResponse.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimative, ber.TagInteger, messageID, "MessageID"))
	searchResponse := ber.Encode(ber.ClassApplication, ber.TypeConstructed, ldap.ApplicationSearchResultEntry, nil, "Search Result Entry")
	searchResponse.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimative, ber.TagOctetString, "", "objectName	LDAPDN"))
	attributesPacket := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Attributes")
	attributesPacket.AppendChild(buildAttributePacket(models.SubschemaSubentryAttribute, cnSchema))
	searchResponse.AppendChild(attributesPacket)
	ldapResponse.AppendChild(searchResponse)

	proc.sendLdapResponse(ldapResponse)

	return ldap.LDAPResultSuccess
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
