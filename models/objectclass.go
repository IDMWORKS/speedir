package models

import "database/sql"

type ObjectClassFlag int

const (
	OCNone = 0
)

const (
	OCStructural ObjectClassFlag = 1 << iota
	OCAuxiliary
	OCAbstract
)

// ObjectClass model in the DB
type ObjectClass struct {
	Name  string
	OID   string
	Super sql.NullString
	Names StringSlice
	Flags ObjectClassFlag

	MustAttributes StringSlice
	MayAttributes  StringSlice
}

func (objectClass *ObjectClass) FlagsString() string {
	switch objectClass.Flags {
	case OCAuxiliary:
		return "AUXILIARY"
	case OCAbstract:
		return "ABSTRACT"
	default:
		return "STRUCTURAL"
	}
}

func (objectClass *ObjectClass) MayString() string {
	return attributesString("MAY", objectClass.MayAttributes)
}

func (objectClass *ObjectClass) MustString() string {
	return attributesString("MUST", objectClass.MustAttributes)
}

func attributesString(prefix string, attributes []string) string {
	attrCount := len(attributes)
	if attrCount == 0 {
		return ""
	}
	result := prefix + " ( "

	last := attrCount - 1
	for i, attr := range attributes {
		result = result + attr
		if i != last {
			result = result + " $ "
		}
	}

	result = result + " )"

	return result
}

const (
	// OIDs
	// https://tools.ietf.org/html/rfc4512
	TopClassID              = "2.5.6.0"
	AliasClassID            = "2.5.6.1"
	SubschemaClassID        = "2.5.20.1"
	ExtensibleObjectClassID = "1.3.6.1.4.1.1466.101.120.111"
	// https://tools.ietf.org/html/rfc4519
	CountryClassID              = "2.5.6.2"
	LocalityClassID             = "2.5.6.3"
	OrganizationClassID         = "2.5.6.4"
	OrganizationalUnitClassID   = "2.5.6.5"
	PersonClassID               = "2.5.6.6"
	OrganizationalPersonClassID = "2.5.6.7"
	OrganizationalRoleClassID   = "2.5.6.8"
	GroupOfNamesClassID         = "2.5.6.9"
	ResidentialPersonClassID    = "2.5.6.10"
	DeviceClassID               = "2.5.6.14"
	ApplicationProcessClassID   = "2.5.6.11"
	GroupOfUniqueNamesClassID   = "2.5.6.17"
	DCObjectClassID             = "1.3.6.1.4.1.1466.344"
	UIDObjectClassID            = "1.3.6.1.1.3.1"

	// names
	// https://tools.ietf.org/html/rfc4512
	TopClass              = "top"
	AliasClass            = "alias"
	SubschemaClass        = "subschema"
	ExtensibleObjectClass = "extensibleObject"
	// https://tools.ietf.org/html/rfc4519
	CountryClass              = "country"
	LocalityClass             = "locality"
	OrganizationClass         = "organization"
	OrganizationalUnitClass   = "organizationalUnit"
	PersonClass               = "person"
	OrganizationalPersonClass = "organizationalPerson"
	OrganizationalRoleClass   = "organizationalRole"
	GroupOfNamesClass         = "groupOfNames"
	ResidentialPersonClass    = "residentialPerson"
	DeviceClass               = "device"
	ApplicationProcessClass   = "applicationProcess"
	GroupOfUniqueNamesClass   = "groupOfUniqueNames"
	DCObjectClass             = "dcObject"
	UIDObjectClass            = "uidObject"
)

// LDAPv3AttributeTypes represents the standard Object Classes
var LDAPv3ObjectClasses = [...]ObjectClass{
	// https://tools.ietf.org/html/rfc4512
	ObjectClass{
		OID:            TopClassID,
		Name:           TopClass,
		MustAttributes: StringSlice{ObjectClassAttribute},
		Flags:          OCAbstract,
	},
	ObjectClass{
		OID:  AliasClassID,
		Name: AliasClass,
		MustAttributes: StringSlice{
			AliasedObjectNameAttribute,
			NameAttribute,
		},
		Flags: OCAbstract,
	},
	ObjectClass{
		OID:  SubschemaClassID,
		Name: SubschemaClass,
		MayAttributes: StringSlice{
			AliasedObjectNameAttribute,
			NameAttribute,
		},
		Flags: OCAuxiliary,
	},
	ObjectClass{
		OID:   ExtensibleObjectClassID,
		Name:  ExtensibleObjectClass,
		Flags: OCAuxiliary,
	},

	// https://tools.ietf.org/html/rfc4519
	ObjectClass{
		OID:            ApplicationProcessClassID,
		Name:           ApplicationProcessClass,
		MustAttributes: StringSlice{CountryNameAttribute},
		MayAttributes: StringSlice{
			SeeAlsoAttribute,
			OrganizationUnitNameAttribute,
			LocalityNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            CountryClassID,
		Name:           CountryClass,
		MustAttributes: StringSlice{CommonNameAttribute},
		MayAttributes: StringSlice{
			SearchGuideAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            DCObjectClassID,
		Name:           DCObjectClass,
		MustAttributes: StringSlice{DomainComponentAttribute},
		Flags:          OCAuxiliary,
	},
	ObjectClass{
		OID:            DeviceClassID,
		Name:           DeviceClass,
		MustAttributes: StringSlice{CommonNameAttribute},
		MayAttributes: StringSlice{
			SerialNumberAttribute,
			SeeAlsoAttribute,
			OwnerAttribute,
			OrganizationUnitNameAttribute,
			OrganizationNameAttribute,
			LocalityNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:  GroupOfNamesClassID,
		Name: GroupOfNamesClass,
		MustAttributes: StringSlice{
			MemberAttribute,
			CommonNameAttribute,
		},
		MayAttributes: StringSlice{
			BusinessCategoryAttribute,
			SeeAlsoAttribute,
			OwnerAttribute,
			OrganizationUnitNameAttribute,
			OrganizationNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:  GroupOfUniqueNamesClassID,
		Name: GroupOfUniqueNamesClass,
		MustAttributes: StringSlice{
			UniqueMemberAttribute,
			CommonNameAttribute,
		},
		MayAttributes: StringSlice{
			BusinessCategoryAttribute,
			SeeAlsoAttribute,
			OwnerAttribute,
			OrganizationUnitNameAttribute,
			OrganizationNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:  LocalityClassID,
		Name: LocalityClass,
		MayAttributes: StringSlice{
			StreetAddressAttribute,
			SeeAlsoAttribute,
			SearchGuideAttribute,
			StateOrProvinceNameAttribute,
			LocalityNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:  OrganizationClassID,
		Name: OrganizationClass,
		MayAttributes: StringSlice{
			UserPasswordAttribute,
			SearchGuideAttribute,
			SeeAlsoAttribute,
			BusinessCategoryAttribute,
			X121AddressAttribute,
			RegisteredAddressAttribute,
			DestinationIndicatorAttribute,
			PreferredDeliveryMethodAttribute,
			TelexNumberAttribute,
			TeletexTerminalIdentAttribute,
			TelephoneNumberAttribute,
			InternationalISDNNumberAttribute,
			FacsimileTelephoneNumberAttribute,
			StreetAddressAttribute,
			PostOfficeBoxAttribute,
			PostalCodeAttribute,
			PostalAddressAttribute,
			PhysicalDeliveryOfficeNameAttribute,
			StateOrProvinceNameAttribute,
			LocalityNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:  PersonClassID,
		Name: PersonClass,
		MustAttributes: StringSlice{
			SurnameAttribute,
			CommonNameAttribute,
		},
		MayAttributes: StringSlice{
			UserPasswordAttribute,
			TelephoneNumberAttribute,
			SeeAlsoAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:   OrganizationalPersonClassID,
		Super: sql.NullString{String: PersonClass, Valid: true},
		Name:  OrganizationalPersonClass,
		MayAttributes: StringSlice{
			TitleAttribute,
			X121AddressAttribute,
			RegisteredAddressAttribute,
			DestinationIndicatorAttribute,
			PreferredDeliveryMethodAttribute,
			TelexNumberAttribute,
			TeletexTerminalIdentAttribute,
			TelephoneNumberAttribute,
			InternationalISDNNumberAttribute,
			FacsimileTelephoneNumberAttribute,
			StreetAddressAttribute,
			PostOfficeBoxAttribute,
			PostalCodeAttribute,
			PostalAddressAttribute,
			PhysicalDeliveryOfficeNameAttribute,
			OrganizationUnitNameAttribute,
			StateOrProvinceNameAttribute,
			LocalityNameAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            OrganizationalRoleClassID,
		Name:           OrganizationalRoleClass,
		MustAttributes: StringSlice{CommonNameAttribute},
		MayAttributes: StringSlice{
			X121AddressAttribute,
			RegisteredAddressAttribute,
			DestinationIndicatorAttribute,
			PreferredDeliveryMethodAttribute,
			TelexNumberAttribute,
			TeletexTerminalIdentAttribute,
			TelephoneNumberAttribute,
			InternationalISDNNumberAttribute,
			FacsimileTelephoneNumberAttribute,
			SeeAlsoAttribute,
			RoleOccupantAttribute,
			PreferredDeliveryMethodAttribute,
			StreetAddressAttribute,
			PostOfficeBoxAttribute,
			PostalCodeAttribute,
			PostalAddressAttribute,
			PhysicalDeliveryOfficeNameAttribute,
			OrganizationUnitNameAttribute,
			StateOrProvinceNameAttribute,
			LocalityNameAttribute,
			DescriptionAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            OrganizationalUnitClassID,
		Name:           OrganizationalUnitClass,
		MustAttributes: StringSlice{OrganizationUnitNameAttribute},
		MayAttributes: StringSlice{
			BusinessCategoryAttribute,
			DescriptionAttribute,
			DestinationIndicatorAttribute,
			FacsimileTelephoneNumberAttribute,
			InternationalISDNNumberAttribute,
			LocalityNameAttribute,
			PhysicalDeliveryOfficeNameAttribute,
			PostalAddressAttribute,
			PostalCodeAttribute,
			PostOfficeBoxAttribute,
			PreferredDeliveryMethodAttribute,
			RegisteredAddressAttribute,
			SearchGuideAttribute,
			SeeAlsoAttribute,
			StateOrProvinceNameAttribute,
			StreetAddressAttribute,
			TelephoneNumberAttribute,
			TeletexTerminalIdentAttribute,
			TelexNumberAttribute,
			UserPasswordAttribute,
			X121AddressAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            ResidentialPersonClassID,
		Super:          sql.NullString{String: PersonClass, Valid: true},
		Name:           ResidentialPersonClass,
		MustAttributes: StringSlice{LocalityNameAttribute},
		MayAttributes: StringSlice{
			BusinessCategoryAttribute,
			X121AddressAttribute,
			RegisteredAddressAttribute,
			DestinationIndicatorAttribute,
			PreferredDeliveryMethodAttribute,
			TelexNumberAttribute,
			TeletexTerminalIdentAttribute,
			TelephoneNumberAttribute,
			InternationalISDNNumberAttribute,
			FacsimileTelephoneNumberAttribute,
			PreferredDeliveryMethodAttribute,
			StreetAddressAttribute,
			PostOfficeBoxAttribute,
			PostalCodeAttribute,
			PostalAddressAttribute,
			PhysicalDeliveryOfficeNameAttribute,
			StateOrProvinceNameAttribute,
			LocalityNameAttribute,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            UIDObjectClassID,
		Name:           UIDObjectClass,
		MustAttributes: StringSlice{UIDAttribute},
		Flags:          OCAuxiliary,
	},
}
