package models

// Syntax represents the LDAPv3 Syntax model
type Syntax struct {
	OID         string
	Description string
}

const (
	AttributeTypeDescriptionSyntaxID = "1.3.6.1.4.1.1466.115.121.1.3"
	BinarySyntaxID                   = "1.3.6.1.4.1.1466.115.121.1.5"
	BitStringSyntaxID                = "1.3.6.1.4.1.1466.115.121.1.6"
	BooleanSyntaxID                  = "1.3.6.1.4.1.1466.115.121.1.7"
	CertificateSyntaxID              = "1.3.6.1.4.1.1466.115.121.1.8"
	CertificateListSyntaxID          = "1.3.6.1.4.1.1466.115.121.1.9"
	CertificatePairSyntaxID          = "1.3.6.1.4.1.1466.115.121.1.10"
	CountryStringSyntaxID            = "1.3.6.1.4.1.1466.115.121.1.11"
	DistinguishedNameSyntaxID        = "1.3.6.1.4.1.1466.115.121.1.12"
	DeliveryMethodSyntaxID           = "1.3.6.1.4.1.1466.115.121.1.14"
	DirectoryStringSyntaxID          = "1.3.6.1.4.1.1466.115.121.1.15"
	DITContentRuleDescSyntaxID       = "1.3.6.1.4.1.1466.115.121.1.16"
	DITStructureRuleDescSyntaxID     = "1.3.6.1.4.1.1466.115.121.1.17"
	EnhancementGuideSyntaxID         = "1.3.6.1.4.1.1466.115.121.1.21"
	FacsimileTelephoneNumberSyntaxID = "1.3.6.1.4.1.1466.115.121.1.22"
	FaxImageSyntaxID                 = "1.3.6.1.4.1.1466.115.121.1.23"
	GeneralizedTimeSyntaxID          = "1.3.6.1.4.1.1466.115.121.1.24"
	GuideSyntaxID                    = "1.3.6.1.4.1.1466.115.121.1.25"
	IA5StringSyntaxID                = "1.3.6.1.4.1.1466.115.121.1.26"
	IntegerSyntaxID                  = "1.3.6.1.4.1.1466.115.121.1.27"
	JPEGImageSyntaxID                = "1.3.6.1.4.1.1466.115.121.1.28"
	MatchingRuleDescSyntaxID         = "1.3.6.1.4.1.1466.115.121.1.30"
	MatchingRuleUseDescSyntaxID      = "1.3.6.1.4.1.1466.115.121.1.31"
	MHSOrAddressSyntaxID             = "1.3.6.1.4.1.1466.115.121.1.33"
	NameAndOptionalUIDSyntaxID       = "1.3.6.1.4.1.1466.115.121.1.34"
	NameFormSyntaxID                 = "1.3.6.1.4.1.1466.115.121.1.35"
	NumericStringSyntaxID            = "1.3.6.1.4.1.1466.115.121.1.36"
	ObjectClassDescriptionSyntaxID   = "1.3.6.1.4.1.1466.115.121.1.37"
	OIDSyntaxID                      = "1.3.6.1.4.1.1466.115.121.1.38"
	OtherMailboxSyntaxID             = "1.3.6.1.4.1.1466.115.121.1.39"
	OctetStringSyntaxID              = "1.3.6.1.4.1.1466.115.121.1.40"
	PostalAddressSyntaxID            = "1.3.6.1.4.1.1466.115.121.1.41"
	PresentationAddressSyntaxID      = "1.3.6.1.4.1.1466.115.121.1.43"
	PrintableStringSyntaxID          = "1.3.6.1.4.1.1466.115.121.1.44"
	SupportedAlgorithmSyntaxID       = "1.3.6.1.4.1.1466.115.121.1.49"
	TelephoneNumberSyntaxID          = "1.3.6.1.4.1.1466.115.121.1.50"
	TeletexTerminalIdentSyntaxID     = "1.3.6.1.4.1.1466.115.121.1.51"
	TelexNumberSyntaxID              = "1.3.6.1.4.1.1466.115.121.1.52"
	UTCTimeSyntaxID                  = "1.3.6.1.4.1.1466.115.121.1.53"
	LDAPSyntaxDescriptionSyntaxID    = "1.3.6.1.4.1.1466.115.121.1.54"
	SubstringAssertionSyntaxID       = "1.3.6.1.4.1.1466.115.121.1.58"
)

// LDAPv3Syntaxes represents the standard LDAPv3 Syntaxes
// http://www.ietf.org/rfc/rfc4517.txt
var LDAPv3Syntaxes = [...]Syntax{
	Syntax{
		OID:         AttributeTypeDescriptionSyntaxID,
		Description: "Attribute Type Description",
	},
	Syntax{
		OID:         BinarySyntaxID,
		Description: "Binary",
	},
	Syntax{
		OID:         BitStringSyntaxID,
		Description: "Bit String",
	},
	Syntax{
		OID:         BooleanSyntaxID,
		Description: "Boolean",
	},
	Syntax{
		OID:         CertificateSyntaxID,
		Description: "Certificate",
	},
	Syntax{
		OID:         CertificateListSyntaxID,
		Description: "Certificate List",
	},
	Syntax{
		OID:         CertificatePairSyntaxID,
		Description: "Certificate Pair",
	},
	Syntax{
		OID:         CountryStringSyntaxID,
		Description: "Country String",
	},
	Syntax{
		OID:         DistinguishedNameSyntaxID,
		Description: "Distinguished Name",
	},
	Syntax{
		OID:         DeliveryMethodSyntaxID,
		Description: "Delivery Method",
	},
	Syntax{
		OID:         DirectoryStringSyntaxID,
		Description: "Directory String",
	},
	Syntax{
		OID:         DITContentRuleDescSyntaxID,
		Description: "DIT Content Rule Description",
	},
	Syntax{
		OID:         DITStructureRuleDescSyntaxID,
		Description: "DIT Structure Rule Description",
	},
	Syntax{
		OID:         EnhancementGuideSyntaxID,
		Description: "Enhanced Guide",
	},
	Syntax{
		OID:         FacsimileTelephoneNumberSyntaxID,
		Description: "Facsimile Telephone Number",
	},
	Syntax{
		OID:         FaxImageSyntaxID,
		Description: "Fax Image",
	},
	Syntax{
		OID:         GeneralizedTimeSyntaxID,
		Description: "Generalized Time",
	},
	Syntax{
		OID:         GuideSyntaxID,
		Description: "Guide",
	},
	Syntax{
		OID:         IA5StringSyntaxID,
		Description: "IA5 String",
	},
	Syntax{
		OID:         IntegerSyntaxID,
		Description: "Integer",
	},
	Syntax{
		OID:         JPEGImageSyntaxID,
		Description: "JPEG Image",
	},
	Syntax{
		OID:         MatchingRuleDescSyntaxID,
		Description: "Matching Rule Description",
	},
	Syntax{
		OID:         MatchingRuleUseDescSyntaxID,
		Description: "Matching Rule Use Description",
	},
	Syntax{
		OID:         MHSOrAddressSyntaxID,
		Description: "MHS OR Address",
	},
	Syntax{
		OID:         NameAndOptionalUIDSyntaxID,
		Description: "Name and Optional UID",
	},
	Syntax{
		OID:         NameFormSyntaxID,
		Description: "Name Form",
	},
	Syntax{
		OID:         NumericStringSyntaxID,
		Description: "Numeric String",
	},
	Syntax{
		OID:         ObjectClassDescriptionSyntaxID,
		Description: "Object Class Description",
	},
	Syntax{
		OID:         OIDSyntaxID,
		Description: "OID",
	},
	Syntax{
		OID:         OtherMailboxSyntaxID,
		Description: "Other Mailbox",
	},
	Syntax{
		OID:         OctetStringSyntaxID,
		Description: "Octet String",
	},
	Syntax{
		OID:         PostalAddressSyntaxID,
		Description: "Postal Address",
	},
	Syntax{
		OID:         PresentationAddressSyntaxID,
		Description: "Presentation Address",
	},
	Syntax{
		OID:         PrintableStringSyntaxID,
		Description: "Printable String",
	},
	Syntax{
		OID:         SupportedAlgorithmSyntaxID,
		Description: "Supported Algorithm",
	},
	Syntax{
		OID:         TelephoneNumberSyntaxID,
		Description: "Telephone Number",
	},
	Syntax{
		OID:         TeletexTerminalIdentSyntaxID,
		Description: "Teletex Terminal Identifier",
	},
	Syntax{
		OID:         TelexNumberSyntaxID,
		Description: "Telex Number",
	},
	Syntax{
		OID:         UTCTimeSyntaxID,
		Description: "UTC Time",
	},
	Syntax{
		OID:         LDAPSyntaxDescriptionSyntaxID,
		Description: "LDAP Syntax Description",
	},
	Syntax{
		OID:         SubstringAssertionSyntaxID,
		Description: "Substring Assertion",
	},
}
