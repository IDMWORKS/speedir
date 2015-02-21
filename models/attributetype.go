package models

import "database/sql"

type AttributeTypeFlag int

const (
	ATNone = 0
)

const (
	ATSingleValue AttributeTypeFlag = 1 << iota
)

// AttributeType model in the DB
type AttributeType struct {
	OID      string
	SuperID  sql.NullString
	SyntaxID sql.NullString
	Names    StringArray
	Flags    AttributeTypeFlag

	EqualityMatchID sql.NullString
	SubstrMatchID   sql.NullString
	OrderingMatchID sql.NullString
}

const (
	CommonNameAttributeID                 = "2.5.4.3"
	SurnameAttributeID                    = "2.5.4.4"
	SerialNumberAttributeID               = "2.5.4.5"
	CountryNameAttributeID                = "2.5.4.6"
	LocalityNameAttributeID               = "2.5.4.7"
	StateOrProvinceNameAttributeID        = "2.5.4.8"
	StreetAddressAttributeID              = "2.5.4.9"
	OrganizationNameAttributeID           = "2.5.4.10"
	OrganizationUnitNameAttributeID       = "2.5.4.11"
	TitleAttributeID                      = "2.5.4.12"
	DescriptionAttributeID                = "2.5.4.13"
	SearchGuideAttributeID                = "2.5.4.14"
	BusinessCategoryAttributeID           = "2.5.4.15"
	PostalAddressAttributeID              = "2.5.4.16"
	PostalCodeAttributeID                 = "2.5.4.17"
	PostOfficeBoxAttributeID              = "2.5.4.18"
	PhysicalDeliveryOfficeNameAttributeID = "2.5.4.19"
	TelephoneNumberAttributeID            = "2.5.4.20"
	TelexNumberAttributeID                = "2.5.4.21"
	TeletexTerminalIdentAttributeID       = "2.5.4.22"
	FacsimileTelephoneNumberAttributeID   = "2.5.4.23"
	X121AddressAttributeID                = "2.5.4.24"
	InternationalISDNNumberAttributeID    = "2.5.4.25"
	RegisteredAddressAttributeID          = "2.5.4.26"
	DestinationIndicatorAttributeID       = "2.5.4.27"
	PreferredDeliveryMethodAttributeID    = "2.5.4.28"
	MemberAttributeID                     = "2.5.4.31"
	OwnerAttributeID                      = "2.5.4.32"
	RoleOccupantAttributeID               = "2.5.4.33"
	SeeAlsoAttributeID                    = "2.5.4.34"
	UserPasswordAttributeID               = "2.5.4.35"
	NameAttributeID                       = "2.5.4.41"
	GivenNameAttributeID                  = "2.5.4.42"
	InitialsAttributeID                   = "2.5.4.43"
	GenerationQualifierAttributeID        = "2.5.4.44"
	X500UniqueIdentifierAttributeID       = "2.5.4.45"
	DNQualifierAttributeID                = "2.5.4.46"
	EnhancedSearchGuideAttributeID        = "2.5.4.47"
	DistinguishedNameAttributeID          = "2.5.4.49"
	UniqueMemberAttributeID               = "2.5.4.50"
	HouseIdentifierAttributeID            = "2.5.4.51"
	DomainComponentAttributeID            = "0.9.2342.19200300.100.1.25"
	UIDAttributeID                        = "0.9.2342.19200300.100.1.1"
)

// LDAPv3AttributeTypes represents the standard Attribute Types
// https://tools.ietf.org/html/rfc4519
var LDAPv3AttributeTypes = [...]AttributeType{
	AttributeType{
		OID:             BusinessCategoryAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"businessCategory"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
		Flags:           ATNone,
	},
	AttributeType{
		OID:             NameAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"name"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:      CountryNameAttributeID,
		SuperID:  sql.NullString{String: NameAttributeID, Valid: true},
		SyntaxID: sql.NullString{String: CountryStringSyntaxID, Valid: true},
		Names:    StringArray{"c", "countryName"},
		Flags:    ATSingleValue,
	},
	AttributeType{
		OID:     CommonNameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"cn", "commonName"},
	},
	AttributeType{
		OID:           DomainComponentAttributeID,
		SyntaxID:      sql.NullString{String: IA5StringSyntaxID, Valid: true},
		Names:         StringArray{"dc", "domainComponent"},
		SubstrMatchID: sql.NullString{String: CaseIgnoreIA5SubstrMatchRuleID, Valid: true},
		Flags:         ATSingleValue,
	},
	AttributeType{
		OID:           DescriptionAttributeID,
		SyntaxID:      sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:         StringArray{"description"},
		SubstrMatchID: sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:           DestinationIndicatorAttributeID,
		SyntaxID:      sql.NullString{String: PrintableStringSyntaxID, Valid: true},
		Names:         StringArray{"destinationIndicator"},
		SubstrMatchID: sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             DistinguishedNameAttributeID,
		SyntaxID:        sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Names:           StringArray{"distinguishedName"},
		EqualityMatchID: sql.NullString{String: DistinguishedNameMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             DNQualifierAttributeID,
		SyntaxID:        sql.NullString{String: PrintableStringSyntaxID, Valid: true},
		Names:           StringArray{"dnQualifier"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		OrderingMatchID: sql.NullString{String: CaseIgnoreOrderingMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:      EnhancedSearchGuideAttributeID,
		SyntaxID: sql.NullString{String: EnhancementGuideSyntaxID, Valid: true},
		Names:    StringArray{"enhancedSearchGuide"},
	},
	AttributeType{
		OID:      FacsimileTelephoneNumberAttributeID,
		SyntaxID: sql.NullString{String: FacsimileTelephoneNumberSyntaxID, Valid: true},
		Names:    StringArray{"facsimileTelephoneNumber"},
	},
	AttributeType{
		OID:     GenerationQualifierAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"generationQualifier"},
	},
	AttributeType{
		OID:     GivenNameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"givenName"},
	},
	AttributeType{
		OID:             HouseIdentifierAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"houseIdentifier"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:     InitialsAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"initials"},
	},
	AttributeType{
		OID:             InternationalISDNNumberAttributeID,
		SyntaxID:        sql.NullString{String: NumericStringSyntaxID, Valid: true},
		Names:           StringArray{"internationalISDNNumber"},
		EqualityMatchID: sql.NullString{String: NumericStringMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: NumericStringSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:     LocalityNameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"l", "localityName"},
	},
	AttributeType{
		OID:     MemberAttributeID,
		SuperID: sql.NullString{String: DistinguishedNameAttributeID, Valid: true},
		Names:   StringArray{"distinguishedName"},
	},
	AttributeType{
		OID:     OrganizationNameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"o", "organizationName"},
	},
	AttributeType{
		OID:     OrganizationUnitNameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"ou", "organizationalUnitName"},
	},
	AttributeType{
		OID:     OwnerAttributeID,
		SuperID: sql.NullString{String: DistinguishedNameAttributeID, Valid: true},
		Names:   StringArray{"owner"},
	},
	AttributeType{
		OID:             PhysicalDeliveryOfficeNameAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"physicalDeliveryOfficeName"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             PostalAddressAttributeID,
		SyntaxID:        sql.NullString{String: PostalAddressSyntaxID, Valid: true},
		Names:           StringArray{"postalAddress"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreListMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreListSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             PostalCodeAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"postalCode"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             PostOfficeBoxAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"postOfficeBox"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:      PreferredDeliveryMethodAttributeID,
		SyntaxID: sql.NullString{String: DeliveryMethodSyntaxID, Valid: true},
		Names:    StringArray{"preferredDeliveryMethod"},
		Flags:    ATSingleValue,
	},
	AttributeType{
		OID:      RegisteredAddressAttributeID,
		SuperID:  sql.NullString{String: PostalAddressAttributeID, Valid: true},
		SyntaxID: sql.NullString{String: PostalAddressSyntaxID, Valid: true},
		Names:    StringArray{"registeredAddress"},
	},
	AttributeType{
		OID:     RoleOccupantAttributeID,
		SuperID: sql.NullString{String: DistinguishedNameAttributeID, Valid: true},
		Names:   StringArray{"roleOccupant"},
	},
	AttributeType{
		OID:      SearchGuideAttributeID,
		SyntaxID: sql.NullString{String: GuideSyntaxID, Valid: true},
		Names:    StringArray{"searchGuide"},
	},
	AttributeType{
		OID:     SeeAlsoAttributeID,
		SuperID: sql.NullString{String: DistinguishedNameAttributeID, Valid: true},
		Names:   StringArray{"seeAlso"},
	},
	AttributeType{
		OID:             SerialNumberAttributeID,
		SyntaxID:        sql.NullString{String: PrintableStringSyntaxID, Valid: true},
		Names:           StringArray{"serialNumber"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:     SurnameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"sn", "surname"},
	},
	AttributeType{
		OID:     StateOrProvinceNameAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"st", "stateOrProvinceName"},
	},
	AttributeType{
		OID:             StreetAddressAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"street", "streetAddress"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             TelephoneNumberAttributeID,
		SyntaxID:        sql.NullString{String: TelephoneNumberSyntaxID, Valid: true},
		Names:           StringArray{"telephoneNumber"},
		EqualityMatchID: sql.NullString{String: TelephoneNumberMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: TelephoneNumberSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:      TeletexTerminalIdentAttributeID,
		SyntaxID: sql.NullString{String: TeletexTerminalIdentSyntaxID, Valid: true},
		Names:    StringArray{"teletexTerminalIdentifier"},
	},
	AttributeType{
		OID:      TelexNumberAttributeID,
		SyntaxID: sql.NullString{String: TelexNumberSyntaxID, Valid: true},
		Names:    StringArray{"telexNumber"},
	},
	AttributeType{
		OID:     TitleAttributeID,
		SuperID: sql.NullString{String: NameAttributeID, Valid: true},
		Names:   StringArray{"title"},
	},
	AttributeType{
		OID:             UIDAttributeID,
		SyntaxID:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Names:           StringArray{"uid"},
		EqualityMatchID: sql.NullString{String: CaseIgnoreMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: CaseIgnoreSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             UniqueMemberAttributeID,
		SyntaxID:        sql.NullString{String: NameAndOptionalUIDSyntaxID, Valid: true},
		Names:           StringArray{"uniqueMember"},
		EqualityMatchID: sql.NullString{String: UniqueMemberMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             UserPasswordAttributeID,
		SyntaxID:        sql.NullString{String: OctetStringSyntaxID, Valid: true},
		Names:           StringArray{"userPassword"},
		EqualityMatchID: sql.NullString{String: OctetStringMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             X121AddressAttributeID,
		SyntaxID:        sql.NullString{String: NumericStringSyntaxID, Valid: true},
		Names:           StringArray{"x121Address"},
		EqualityMatchID: sql.NullString{String: NumericStringMatchRuleID, Valid: true},
		SubstrMatchID:   sql.NullString{String: NumericStringSubstrMatchRuleID, Valid: true},
	},
	AttributeType{
		OID:             X500UniqueIdentifierAttributeID,
		SyntaxID:        sql.NullString{String: BitStringSyntaxID, Valid: true},
		Names:           StringArray{"x500UniqueIdentifier"},
		EqualityMatchID: sql.NullString{String: BitStringMatchRuleID, Valid: true},
	},
}
