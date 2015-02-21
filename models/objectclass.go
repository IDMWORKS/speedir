package models

import "database/sql"

type ObjectClassFlag int

const (
	OCNone = 0
)

const (
	OCStructural ObjectClassFlag = 1 << iota
	OCAuxiliary
)

// ObjectClass model in the DB
type ObjectClass struct {
	OID     string
	SuperID sql.NullString
	Names   StringArray
	Flags   ObjectClassFlag

	MustAttributes StringArray
	MayAttributes  StringArray
}

const (
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
	GroupOfUniqueNamesID        = "2.5.6.17"
	DCObjectClassID             = "1.3.6.1.4.1.1466.344"
	UIDObjectClassID            = "1.3.6.1.1.3.1"
)

// LDAPv3AttributeTypes represents the standard Object Classes
// https://tools.ietf.org/html/rfc4519
var LDAPv3ObjectClasses = [...]ObjectClass{
	ObjectClass{
		OID:            ApplicationProcessClassID,
		Names:          StringArray{"applicationProcess"},
		MustAttributes: StringArray{CountryNameAttributeID},
		MayAttributes: StringArray{
			SeeAlsoAttributeID,
			OrganizationUnitNameAttributeID,
			LocalityNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            CountryClassID,
		Names:          StringArray{"country"},
		MustAttributes: StringArray{CommonNameAttributeID},
		MayAttributes: StringArray{
			SearchGuideAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            DCObjectClassID,
		Names:          StringArray{"dcObject"},
		MustAttributes: StringArray{DomainComponentAttributeID},
		MayAttributes:  StringArray{},
		Flags:          OCAuxiliary,
	},
	ObjectClass{
		OID:            DeviceClassID,
		Names:          StringArray{"device"},
		MustAttributes: StringArray{CommonNameAttributeID},
		MayAttributes: StringArray{
			SerialNumberAttributeID,
			SeeAlsoAttributeID,
			OwnerAttributeID,
			OrganizationUnitNameAttributeID,
			OrganizationNameAttributeID,
			LocalityNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:   GroupOfNamesClassID,
		Names: StringArray{"groupOfNames"},
		MustAttributes: StringArray{
			MemberAttributeID,
			CommonNameAttributeID,
		},
		MayAttributes: StringArray{
			BusinessCategoryAttributeID,
			SeeAlsoAttributeID,
			OwnerAttributeID,
			OrganizationUnitNameAttributeID,
			OrganizationNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:   GroupOfUniqueNamesID,
		Names: StringArray{"groupOfUniqueNames"},
		MustAttributes: StringArray{
			UniqueMemberAttributeID,
			CommonNameAttributeID,
		},
		MayAttributes: StringArray{
			BusinessCategoryAttributeID,
			SeeAlsoAttributeID,
			OwnerAttributeID,
			OrganizationUnitNameAttributeID,
			OrganizationNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            LocalityClassID,
		Names:          StringArray{"locality"},
		MustAttributes: StringArray{},
		MayAttributes: StringArray{
			StreetAddressAttributeID,
			SeeAlsoAttributeID,
			SearchGuideAttributeID,
			StateOrProvinceNameAttributeID,
			LocalityNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            OrganizationClassID,
		Names:          StringArray{"organization"},
		MustAttributes: StringArray{},
		MayAttributes: StringArray{
			UserPasswordAttributeID,
			SearchGuideAttributeID,
			SeeAlsoAttributeID,
			BusinessCategoryAttributeID,
			X121AddressAttributeID,
			RegisteredAddressAttributeID,
			DestinationIndicatorAttributeID,
			PreferredDeliveryMethodAttributeID,
			TelexNumberAttributeID,
			TeletexTerminalIdentAttributeID,
			TelephoneNumberAttributeID,
			InternationalISDNNumberAttributeID,
			FacsimileTelephoneNumberAttributeID,
			StreetAddressAttributeID,
			PostOfficeBoxAttributeID,
			PostalCodeAttributeID,
			PostalAddressAttributeID,
			PhysicalDeliveryOfficeNameAttributeID,
			StateOrProvinceNameAttributeID,
			LocalityNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:   PersonClassID,
		Names: StringArray{"person"},
		MustAttributes: StringArray{
			SurnameAttributeID,
			CommonNameAttributeID,
		},
		MayAttributes: StringArray{
			UserPasswordAttributeID,
			TelephoneNumberAttributeID,
			SeeAlsoAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            OrganizationalPersonClassID,
		SuperID:        sql.NullString{String: PersonClassID, Valid: true},
		Names:          StringArray{"organizationalPerson"},
		MustAttributes: StringArray{},
		MayAttributes: StringArray{
			TitleAttributeID,
			X121AddressAttributeID,
			RegisteredAddressAttributeID,
			DestinationIndicatorAttributeID,
			PreferredDeliveryMethodAttributeID,
			TelexNumberAttributeID,
			TeletexTerminalIdentAttributeID,
			TelephoneNumberAttributeID,
			InternationalISDNNumberAttributeID,
			FacsimileTelephoneNumberAttributeID,
			StreetAddressAttributeID,
			PostOfficeBoxAttributeID,
			PostalCodeAttributeID,
			PostalAddressAttributeID,
			PhysicalDeliveryOfficeNameAttributeID,
			OrganizationUnitNameAttributeID,
			StateOrProvinceNameAttributeID,
			LocalityNameAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            OrganizationalRoleClassID,
		Names:          StringArray{"organizationalRole"},
		MustAttributes: StringArray{CommonNameAttributeID},
		MayAttributes: StringArray{
			X121AddressAttributeID,
			RegisteredAddressAttributeID,
			DestinationIndicatorAttributeID,
			PreferredDeliveryMethodAttributeID,
			TelexNumberAttributeID,
			TeletexTerminalIdentAttributeID,
			TelephoneNumberAttributeID,
			InternationalISDNNumberAttributeID,
			FacsimileTelephoneNumberAttributeID,
			SeeAlsoAttributeID,
			RoleOccupantAttributeID,
			PreferredDeliveryMethodAttributeID,
			StreetAddressAttributeID,
			PostOfficeBoxAttributeID,
			PostalCodeAttributeID,
			PostalAddressAttributeID,
			PhysicalDeliveryOfficeNameAttributeID,
			OrganizationUnitNameAttributeID,
			StateOrProvinceNameAttributeID,
			LocalityNameAttributeID,
			DescriptionAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            OrganizationalUnitClassID,
		Names:          StringArray{"organizationalUnit"},
		MustAttributes: StringArray{OrganizationUnitNameAttributeID},
		MayAttributes: StringArray{
			BusinessCategoryAttributeID,
			DescriptionAttributeID,
			DestinationIndicatorAttributeID,
			FacsimileTelephoneNumberAttributeID,
			InternationalISDNNumberAttributeID,
			LocalityNameAttributeID,
			PhysicalDeliveryOfficeNameAttributeID,
			PostalAddressAttributeID,
			PostalCodeAttributeID,
			PostOfficeBoxAttributeID,
			PreferredDeliveryMethodAttributeID,
			RegisteredAddressAttributeID,
			SearchGuideAttributeID,
			SeeAlsoAttributeID,
			StateOrProvinceNameAttributeID,
			StreetAddressAttributeID,
			TelephoneNumberAttributeID,
			TeletexTerminalIdentAttributeID,
			TelexNumberAttributeID,
			UserPasswordAttributeID,
			X121AddressAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            ResidentialPersonClassID,
		SuperID:        sql.NullString{String: PersonClassID, Valid: true},
		Names:          StringArray{"residentialPerson"},
		MustAttributes: StringArray{LocalityNameAttributeID},
		MayAttributes: StringArray{
			BusinessCategoryAttributeID,
			X121AddressAttributeID,
			RegisteredAddressAttributeID,
			DestinationIndicatorAttributeID,
			PreferredDeliveryMethodAttributeID,
			TelexNumberAttributeID,
			TeletexTerminalIdentAttributeID,
			TelephoneNumberAttributeID,
			InternationalISDNNumberAttributeID,
			FacsimileTelephoneNumberAttributeID,
			PreferredDeliveryMethodAttributeID,
			StreetAddressAttributeID,
			PostOfficeBoxAttributeID,
			PostalCodeAttributeID,
			PostalAddressAttributeID,
			PhysicalDeliveryOfficeNameAttributeID,
			StateOrProvinceNameAttributeID,
			LocalityNameAttributeID,
		},
		Flags: OCStructural,
	},
	ObjectClass{
		OID:            UIDObjectClassID,
		Names:          StringArray{"uidObject"},
		MustAttributes: StringArray{UIDAttributeID},
		MayAttributes:  StringArray{},
		Flags:          OCAuxiliary,
	},
}
