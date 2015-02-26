package models

import "database/sql"

type AttributeTypeFlag int
type AttributeUsageFlag int

const (
	ATNone = 0
	AUNone = 0
)

const (
	ATSingleValue AttributeTypeFlag = 1 << iota
	ATNoUserMods
	ATObsolete
	ATCollective
)

const (
	AUDSAOperation AttributeUsageFlag = 1 << iota
	AUDirectoryOperation
	AUUserApplications
	AUDistributedOperation
)

// AttributeType model in the DB
type AttributeType struct {
	Name   string
	OID    string
	Super  sql.NullString
	Syntax sql.NullString
	Names  StringSlice
	Flags  AttributeTypeFlag
	Usage  AttributeUsageFlag

	EqualityMatch sql.NullString
	SubstrMatch   sql.NullString
	OrderingMatch sql.NullString
}

const (
	// OIDs
	// https://tools.ietf.org/html/rfc4512
	ObjectClassAttributeID                 = "2.5.4.0"
	AliasedObjectNameAttributeID           = "2.5.4.1"
	CreateTimestampAttributeID             = "2.5.18.1"
	ModifyTimestampAttributeID             = "2.5.18.2"
	CreatorsNameAttributeID                = "2.5.18.3"
	ModifiersNameAttributeID               = "2.5.18.4"
	SubschemaSubentryAttributeID           = "2.5.18.10"
	DITStructureRulesAttributeID           = "2.5.21.1"
	DITContentRulesAttributeID             = "2.5.21.2"
	MatchingRulesAttributeID               = "2.5.21.4"
	AttributeTypesAttributeID              = "2.5.21.5"
	ObjectClassesAttributeID               = "2.5.21.6"
	NameFormsAttributeID                   = "2.5.21.7"
	MatchingRuleUseAttributeID             = "2.5.21.8"
	SupportedFeaturesAttributeID           = "1.3.6.1.4.1.4203.1.3.5"
	NamingContextsAttributeID              = "1.3.6.1.4.1.1466.101.120.5"
	AltServerAttributeID                   = "1.3.6.1.4.1.1466.101.120.6"
	SupportedExtensionAttributeID          = "1.3.6.1.4.1.1466.101.120.7"
	SupportedControlAttributeID            = "1.3.6.1.4.1.1466.101.120.13"
	SupportedLDAPSASLMechanismsAttributeID = "1.3.6.1.4.1.1466.101.120.14"
	SupportedLDAPVersionAttributeID        = "1.3.6.1.4.1.1466.101.120.15"
	LDAPSyntaxesAttributeID                = "1.3.6.1.4.1.1466.101.120.16"
	// https://tools.ietf.org/html/rfc4519
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

	// names
	// https://tools.ietf.org/html/rfc4512
	ObjectClassAttribute                 = "objectClass"
	AliasedObjectNameAttribute           = "aliasedObjectName"
	CreateTimestampAttribute             = "createTimestamp"
	ModifyTimestampAttribute             = "modifyTimestamp"
	CreatorsNameAttribute                = "creatorsName"
	ModifiersNameAttribute               = "modifiersName"
	SubschemaSubentryAttribute           = "subschemaSubentry"
	DITStructureRulesAttribute           = "dITStructureRules"
	DITContentRulesAttribute             = "dITContentRules"
	MatchingRulesAttribute               = "matchingRules"
	AttributeTypesAttribute              = "attributeTypes"
	ObjectClassesAttribute               = "objectClasses"
	NameFormsAttribute                   = "nameForms"
	MatchingRuleUseAttribute             = "matchingRuleUse"
	SupportedFeaturesAttribute           = "supportedFeatures"
	NamingContextsAttribute              = "namingContexts"
	AltServerAttribute                   = "altServer"
	SupportedExtensionAttribute          = "supportedExtension"
	SupportedControlAttribute            = "supportedControl"
	SupportedLDAPSASLMechanismsAttribute = "supportedSASLMechanisms"
	SupportedLDAPVersionAttribute        = "supportedLDAPVersion"
	LDAPSyntaxesAttribute                = "ldapSyntaxes"
	// https://tools.ietf.org/html/rfc4519
	CommonNameAttribute                 = "cn"
	SurnameAttribute                    = "sn"
	SerialNumberAttribute               = "serialNumber"
	CountryNameAttribute                = "c"
	LocalityNameAttribute               = "l"
	StateOrProvinceNameAttribute        = "st"
	StreetAddressAttribute              = "street"
	OrganizationNameAttribute           = "o"
	OrganizationUnitNameAttribute       = "ou"
	TitleAttribute                      = "title"
	DescriptionAttribute                = "description"
	SearchGuideAttribute                = "searchGuide"
	BusinessCategoryAttribute           = "businessCategory"
	PostalAddressAttribute              = "postalAddress"
	PostalCodeAttribute                 = "postalCode"
	PostOfficeBoxAttribute              = "postOfficeBox"
	PhysicalDeliveryOfficeNameAttribute = "physicalDeliveryOfficeName"
	TelephoneNumberAttribute            = "telephoneNumber"
	TelexNumberAttribute                = "telexNumber"
	TeletexTerminalIdentAttribute       = "teletexTerminalIdentifier"
	FacsimileTelephoneNumberAttribute   = "facsimileTelephoneNumber"
	X121AddressAttribute                = "x121Address"
	InternationalISDNNumberAttribute    = "internationalISDNNumber"
	RegisteredAddressAttribute          = "registeredAddress"
	DestinationIndicatorAttribute       = "destinationIndicator"
	PreferredDeliveryMethodAttribute    = "preferredDeliveryMethod"
	MemberAttribute                     = "member"
	OwnerAttribute                      = "owner"
	RoleOccupantAttribute               = "roleOccupant"
	SeeAlsoAttribute                    = "seeAlso"
	UserPasswordAttribute               = "userPassword"
	NameAttribute                       = "name"
	GivenNameAttribute                  = "givenName"
	InitialsAttribute                   = "initials"
	GenerationQualifierAttribute        = "generationQualifier"
	X500UniqueIdentifierAttribute       = "x500UniqueIdentifier"
	DNQualifierAttribute                = "dnQualifier"
	EnhancedSearchGuideAttribute        = "enhancedSearchGuide"
	DistinguishedNameAttribute          = "distinguishedName"
	UniqueMemberAttribute               = "uniqueMember"
	HouseIdentifierAttribute            = "houseIdentifier"
	DomainComponentAttribute            = "dc"
	UIDAttribute                        = "uid"
)

// LDAPv3AttributeTypes represents the standard Attribute Types
var LDAPv3AttributeTypes = [...]AttributeType{
	// https://tools.ietf.org/html/rfc4512
	AttributeType{
		OID:           ObjectClassAttributeID,
		Syntax:        sql.NullString{String: OIDSyntaxID, Valid: true},
		Name:          ObjectClassAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierMatchRule, Valid: true},
		Flags:         ATNone,
	},
	AttributeType{
		OID:           AliasedObjectNameAttributeID,
		Syntax:        sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Name:          AliasedObjectNameAttribute,
		EqualityMatch: sql.NullString{String: DistinguishedNameMatchRule, Valid: true},
		Flags:         ATSingleValue,
	},
	AttributeType{
		OID:           CreateTimestampAttributeID,
		Syntax:        sql.NullString{String: GeneralizedTimeSyntaxID, Valid: true},
		Name:          CreateTimestampAttribute,
		EqualityMatch: sql.NullString{String: GeneralizedTimeMatchRule, Valid: true},
		OrderingMatch: sql.NullString{String: GeneralizedTimeOrderingMatchRule, Valid: true},
		Flags:         ATSingleValue | ATNoUserMods,
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           ModifyTimestampAttributeID,
		Syntax:        sql.NullString{String: GeneralizedTimeSyntaxID, Valid: true},
		Name:          ModifyTimestampAttribute,
		EqualityMatch: sql.NullString{String: GeneralizedTimeMatchRule, Valid: true},
		OrderingMatch: sql.NullString{String: GeneralizedTimeOrderingMatchRule, Valid: true},
		Flags:         ATSingleValue | ATNoUserMods,
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           CreatorsNameAttributeID,
		Syntax:        sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Name:          CreatorsNameAttribute,
		EqualityMatch: sql.NullString{String: DistinguishedNameMatchRule, Valid: true},
		Flags:         ATSingleValue | ATNoUserMods,
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           ModifiersNameAttributeID,
		Syntax:        sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Name:          ModifiersNameAttribute,
		EqualityMatch: sql.NullString{String: DistinguishedNameMatchRule, Valid: true},
		Flags:         ATSingleValue | ATNoUserMods,
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           SubschemaSubentryAttributeID,
		Syntax:        sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Name:          SubschemaSubentryAttribute,
		EqualityMatch: sql.NullString{String: DistinguishedNameMatchRule, Valid: true},
		Flags:         ATSingleValue | ATNoUserMods,
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           DITStructureRulesAttributeID,
		Syntax:        sql.NullString{String: DITStructureRuleDescSyntaxID, Valid: true},
		Name:          DITStructureRulesAttribute,
		EqualityMatch: sql.NullString{String: IntegerFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           DITContentRulesAttributeID,
		Syntax:        sql.NullString{String: DITContentRuleDescSyntaxID, Valid: true},
		Name:          DITContentRulesAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           MatchingRulesAttributeID,
		Syntax:        sql.NullString{String: MatchingRuleDescSyntaxID, Valid: true},
		Name:          MatchingRulesAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           AttributeTypesAttributeID,
		Syntax:        sql.NullString{String: AttributeTypeDescriptionSyntaxID, Valid: true},
		Name:          AttributeTypesAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           ObjectClassesAttributeID,
		Syntax:        sql.NullString{String: ObjectClassDescriptionSyntaxID, Valid: true},
		Name:          ObjectClassesAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           NameFormsAttributeID,
		Syntax:        sql.NullString{String: NameFormSyntaxID, Valid: true},
		Name:          NameFormsAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           MatchingRuleUseAttributeID,
		Syntax:        sql.NullString{String: MatchingRuleUseDescSyntaxID, Valid: true},
		Name:          MatchingRuleUseAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierFirstCompMatchRule, Valid: true},
		Usage:         AUDirectoryOperation,
	},
	AttributeType{
		OID:           SupportedFeaturesAttributeID,
		Syntax:        sql.NullString{String: OIDSyntaxID, Valid: true},
		Name:          SupportedFeaturesAttribute,
		EqualityMatch: sql.NullString{String: ObjectIdentifierMatchRule, Valid: true},
		Usage:         AUDSAOperation,
	},
	AttributeType{
		OID:    NamingContextsAttributeID,
		Syntax: sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Name:   NamingContextsAttribute,
		Usage:  AUDSAOperation,
	},
	AttributeType{
		OID:    AltServerAttributeID,
		Syntax: sql.NullString{String: IA5StringSyntaxID, Valid: true},
		Name:   AltServerAttribute,
		Usage:  AUDSAOperation,
	},
	AttributeType{
		OID:    SupportedExtensionAttributeID,
		Syntax: sql.NullString{String: OIDSyntaxID, Valid: true},
		Name:   SupportedExtensionAttribute,
		Usage:  AUDSAOperation,
	},
	AttributeType{
		OID:    SupportedControlAttributeID,
		Syntax: sql.NullString{String: OIDSyntaxID, Valid: true},
		Name:   SupportedControlAttribute,
		Usage:  AUDSAOperation,
	},
	AttributeType{
		OID:    SupportedLDAPSASLMechanismsAttributeID,
		Syntax: sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:   SupportedLDAPSASLMechanismsAttribute,
		Usage:  AUDSAOperation,
	},
	AttributeType{
		OID:    SupportedLDAPVersionAttributeID,
		Syntax: sql.NullString{String: IntegerSyntaxID, Valid: true},
		Name:   SupportedLDAPVersionAttribute,
		Usage:  AUDSAOperation,
	},
	AttributeType{
		OID:    LDAPSyntaxesAttributeID,
		Syntax: sql.NullString{String: LDAPSyntaxDescriptionSyntaxID, Valid: true},
		Name:   LDAPSyntaxesAttribute,
		Usage:  AUDirectoryOperation,
	},

	// https://tools.ietf.org/html/rfc4519
	AttributeType{
		OID:           BusinessCategoryAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          BusinessCategoryAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
		Flags:         ATNone,
	},
	AttributeType{
		OID:           NameAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          NameAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:    CountryNameAttributeID,
		Super:  sql.NullString{String: NameAttribute, Valid: true},
		Syntax: sql.NullString{String: CountryStringSyntaxID, Valid: true},
		Name:   CountryNameAttribute,
		Names:  StringSlice{"countryName"},
		Flags:  ATSingleValue,
	},
	AttributeType{
		OID:   CommonNameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  CommonNameAttribute,
		Names: StringSlice{"commonName"},
	},
	AttributeType{
		OID:         DomainComponentAttributeID,
		Syntax:      sql.NullString{String: IA5StringSyntaxID, Valid: true},
		Name:        DomainComponentAttribute,
		Names:       StringSlice{"domainComponent"},
		SubstrMatch: sql.NullString{String: CaseIgnoreIA5SubstrMatchRule, Valid: true},
		Flags:       ATSingleValue,
	},
	AttributeType{
		OID:         DescriptionAttributeID,
		Syntax:      sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:        DescriptionAttribute,
		SubstrMatch: sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:         DestinationIndicatorAttributeID,
		Syntax:      sql.NullString{String: PrintableStringSyntaxID, Valid: true},
		Name:        DestinationIndicatorAttribute,
		SubstrMatch: sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           DistinguishedNameAttributeID,
		Syntax:        sql.NullString{String: DistinguishedNameSyntaxID, Valid: true},
		Name:          DistinguishedNameAttribute,
		EqualityMatch: sql.NullString{String: DistinguishedNameMatchRule, Valid: true},
	},
	AttributeType{
		OID:           DNQualifierAttributeID,
		Syntax:        sql.NullString{String: PrintableStringSyntaxID, Valid: true},
		Name:          DNQualifierAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		OrderingMatch: sql.NullString{String: CaseIgnoreOrderingMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:    EnhancedSearchGuideAttributeID,
		Syntax: sql.NullString{String: EnhancementGuideSyntaxID, Valid: true},
		Name:   EnhancedSearchGuideAttribute,
	},
	AttributeType{
		OID:    FacsimileTelephoneNumberAttributeID,
		Syntax: sql.NullString{String: FacsimileTelephoneNumberSyntaxID, Valid: true},
		Name:   FacsimileTelephoneNumberAttribute,
	},
	AttributeType{
		OID:   GenerationQualifierAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  GenerationQualifierAttribute,
	},
	AttributeType{
		OID:   GivenNameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  GivenNameAttribute,
	},
	AttributeType{
		OID:           HouseIdentifierAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          HouseIdentifierAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:   InitialsAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  InitialsAttribute,
	},
	AttributeType{
		OID:           InternationalISDNNumberAttributeID,
		Syntax:        sql.NullString{String: NumericStringSyntaxID, Valid: true},
		Name:          InternationalISDNNumberAttribute,
		EqualityMatch: sql.NullString{String: NumericStringMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: NumericStringSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:   LocalityNameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  LocalityNameAttribute,
		Names: StringSlice{"localityName"},
	},
	AttributeType{
		OID:   MemberAttributeID,
		Super: sql.NullString{String: DistinguishedNameAttribute, Valid: true},
		Name:  MemberAttribute,
	},
	AttributeType{
		OID:   OrganizationNameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  OrganizationNameAttribute,
		Names: StringSlice{"organizationName"},
	},
	AttributeType{
		OID:   OrganizationUnitNameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  OrganizationUnitNameAttribute,
		Names: StringSlice{"organizationalUnitName"},
	},
	AttributeType{
		OID:   OwnerAttributeID,
		Super: sql.NullString{String: DistinguishedNameAttribute, Valid: true},
		Name:  OwnerAttribute,
	},
	AttributeType{
		OID:           PhysicalDeliveryOfficeNameAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          PhysicalDeliveryOfficeNameAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           PostalAddressAttributeID,
		Syntax:        sql.NullString{String: PostalAddressSyntaxID, Valid: true},
		Name:          PostalAddressAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreListMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreListSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           PostalCodeAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          PostalCodeAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           PostOfficeBoxAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          PostOfficeBoxAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:    PreferredDeliveryMethodAttributeID,
		Syntax: sql.NullString{String: DeliveryMethodSyntaxID, Valid: true},
		Name:   PreferredDeliveryMethodAttribute,
		Flags:  ATSingleValue,
	},
	AttributeType{
		OID:    RegisteredAddressAttributeID,
		Super:  sql.NullString{String: PostalAddressAttribute, Valid: true},
		Syntax: sql.NullString{String: PostalAddressSyntaxID, Valid: true},
		Name:   RegisteredAddressAttribute,
	},
	AttributeType{
		OID:   RoleOccupantAttributeID,
		Super: sql.NullString{String: DistinguishedNameAttribute, Valid: true},
		Name:  RoleOccupantAttribute,
	},
	AttributeType{
		OID:    SearchGuideAttributeID,
		Syntax: sql.NullString{String: GuideSyntaxID, Valid: true},
		Name:   SearchGuideAttribute,
	},
	AttributeType{
		OID:   SeeAlsoAttributeID,
		Super: sql.NullString{String: DistinguishedNameAttribute, Valid: true},
		Name:  SeeAlsoAttribute,
	},
	AttributeType{
		OID:           SerialNumberAttributeID,
		Syntax:        sql.NullString{String: PrintableStringSyntaxID, Valid: true},
		Name:          SerialNumberAttribute,
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:   SurnameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  SurnameAttribute,
		Names: StringSlice{"surname"},
	},
	AttributeType{
		OID:   StateOrProvinceNameAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  StateOrProvinceNameAttribute,
		Names: StringSlice{"stateOrProvinceName"},
	},
	AttributeType{
		OID:           StreetAddressAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          StreetAddressAttribute,
		Names:         StringSlice{"streetAddress"},
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           TelephoneNumberAttributeID,
		Syntax:        sql.NullString{String: TelephoneNumberSyntaxID, Valid: true},
		Name:          TelephoneNumberAttribute,
		EqualityMatch: sql.NullString{String: TelephoneNumberMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: TelephoneNumberSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:    TeletexTerminalIdentAttributeID,
		Syntax: sql.NullString{String: TeletexTerminalIdentSyntaxID, Valid: true},
		Name:   TeletexTerminalIdentAttribute,
	},
	AttributeType{
		OID:    TelexNumberAttributeID,
		Syntax: sql.NullString{String: TelexNumberSyntaxID, Valid: true},
		Name:   TelexNumberAttribute,
	},
	AttributeType{
		OID:   TitleAttributeID,
		Super: sql.NullString{String: NameAttribute, Valid: true},
		Name:  TitleAttribute,
	},
	AttributeType{
		OID:           UIDAttributeID,
		Syntax:        sql.NullString{String: DirectoryStringSyntaxID, Valid: true},
		Name:          UIDAttribute,
		Names:         StringSlice{"userid"},
		EqualityMatch: sql.NullString{String: CaseIgnoreMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: CaseIgnoreSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           UniqueMemberAttributeID,
		Syntax:        sql.NullString{String: NameAndOptionalUIDSyntaxID, Valid: true},
		Name:          UniqueMemberAttribute,
		EqualityMatch: sql.NullString{String: UniqueMemberMatchRule, Valid: true},
	},
	AttributeType{
		OID:           UserPasswordAttributeID,
		Syntax:        sql.NullString{String: OctetStringSyntaxID, Valid: true},
		Name:          UserPasswordAttribute,
		EqualityMatch: sql.NullString{String: OctetStringMatchRule, Valid: true},
	},
	AttributeType{
		OID:           X121AddressAttributeID,
		Syntax:        sql.NullString{String: NumericStringSyntaxID, Valid: true},
		Name:          X121AddressAttribute,
		EqualityMatch: sql.NullString{String: NumericStringMatchRule, Valid: true},
		SubstrMatch:   sql.NullString{String: NumericStringSubstrMatchRule, Valid: true},
	},
	AttributeType{
		OID:           X500UniqueIdentifierAttributeID,
		Syntax:        sql.NullString{String: BitStringSyntaxID, Valid: true},
		Name:          X500UniqueIdentifierAttribute,
		EqualityMatch: sql.NullString{String: BitStringMatchRule, Valid: true},
	},
}
