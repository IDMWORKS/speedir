package models

// MatchingRule model in the DB
type MatchingRule struct {
	OID      string
	SyntaxID string
	Names    StringArray
}

const (
	// ObjectIdentifierMatchRuleID ... http://www.ietf.org/rfc/rfc4517.txt
	ObjectIdentifierMatchRuleID          = "2.5.13.0"
	DistinguishedNameMatchRuleID         = "2.5.13.1"
	CaseIgnoreMatchRuleID                = "2.5.13.2"
	CaseIgnoreOrderingMatchRuleID        = "2.5.13.3"
	CaseIgnoreSubstrMatchRuleID          = "2.5.13.4"
	CaseExactMatchRuleID                 = "2.5.13.5"
	CaseExactOrderingMatchRuleID         = "2.5.13.6"
	CaseExactSubstrMatchRuleID           = "2.5.13.7"
	NumericStringMatchRuleID             = "2.5.13.8"
	NumericStringOrderingMatchRuleID     = "2.5.13.9"
	NumericStringSubstrMatchRuleID       = "2.5.13.10"
	CaseIgnoreListMatchRuleID            = "2.5.13.11"
	CaseIgnoreListSubstrMatchRuleID      = "2.5.13.12"
	BooleanMatchRuleID                   = "2.5.13.13"
	IntegerMatchRuleID                   = "2.5.13.14"
	IntegerOrderingMatchRuleID           = "2.5.13.15"
	BitStringMatchRuleID                 = "2.5.13.16"
	OctetStringMatchRuleID               = "2.5.13.17"
	OctetStringOrderingMatchRuleID       = "2.5.13.18"
	TelephoneNumberMatchRuleID           = "2.5.13.20"
	TelephoneNumberSubstrMatchRuleID     = "2.5.13.21"
	UniqueMemberMatchRuleID              = "2.5.13.23"
	GeneralizedTimeMatchRuleID           = "2.5.13.27"
	GeneralizedTimeOrderingMatchRuleID   = "2.5.13.28"
	IntegerFirstCompMatchRuleID          = "2.5.13.29"
	ObjectIdentifierFirstCompMatchRuleID = "2.5.13.30"
	DirectoryStringFirstCompMatchRuleID  = "2.5.13.31"
	KeywordMatchRuleID                   = "2.5.13.33"
	WordMatchRuleID                      = "2.5.13.32"
	CaseExactIA5MatchRuleID              = "1.3.6.1.4.1.1466.109.114.1"
	CaseIgnoreIA5MatchruleID             = "1.3.6.1.4.1.1466.109.114.2"
	CaseIgnoreIA5SubstrMatchRuleID       = "1.3.6.1.4.1.1466.109.114.3"
)

// LDAPv3MatchingRules represents the standard LDAPv3 Matching Rules
// http://www.ietf.org/rfc/rfc4517.txt
var LDAPv3MatchingRules = [...]MatchingRule{
	MatchingRule{
		OID:      BitStringMatchRuleID,
		SyntaxID: BitStringSyntaxID,
		Names:    StringArray{"bitStringMatch"},
	},
	MatchingRule{
		OID:      BooleanMatchRuleID,
		SyntaxID: BooleanSyntaxID,
		Names:    StringArray{"booleanMatch"},
	},
	MatchingRule{
		OID:      CaseExactIA5MatchRuleID,
		SyntaxID: IA5StringSyntaxID,
		Names:    StringArray{"caseExactIA5Match"},
	},
	MatchingRule{
		OID:      CaseExactMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"caseExactMatch"},
	},
	MatchingRule{
		OID:      CaseExactOrderingMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"caseExactOrderingMatch"},
	},
	MatchingRule{
		OID:      CaseExactSubstrMatchRuleID,
		SyntaxID: SubstringAssertionSyntaxID,
		Names:    StringArray{"caseExactSubstringsMatch"},
	},
	MatchingRule{
		OID:      CaseIgnoreIA5MatchruleID,
		SyntaxID: IA5StringSyntaxID,
		Names:    StringArray{"caseIgnoreIA5Match"},
	},
	MatchingRule{
		OID:      CaseIgnoreIA5SubstrMatchRuleID,
		SyntaxID: SubstringAssertionSyntaxID,
		Names:    StringArray{"caseIgnoreIA5SubstringsMatch"},
	},
	MatchingRule{
		OID:      CaseIgnoreListMatchRuleID,
		SyntaxID: PostalAddressSyntaxID,
		Names:    StringArray{"caseIgnoreListMatch"},
	},
	MatchingRule{
		OID:      CaseIgnoreListSubstrMatchRuleID,
		SyntaxID: SubstringAssertionSyntaxID,
		Names:    StringArray{"caseIgnoreListSubstringsMatch"},
	},
	MatchingRule{
		OID:      CaseIgnoreMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"caseIgnoreMatch"},
	},
	MatchingRule{
		OID:      CaseIgnoreOrderingMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"caseIgnoreOrderingMatch"},
	},
	MatchingRule{
		OID:      CaseIgnoreSubstrMatchRuleID,
		SyntaxID: SubstringAssertionSyntaxID,
		Names:    StringArray{"caseIgnoreSubstringsMatch"},
	},
	MatchingRule{
		OID:      DirectoryStringFirstCompMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"directoryStringFirstComponentMatch"},
	},
	MatchingRule{
		OID:      DistinguishedNameMatchRuleID,
		SyntaxID: DistinguishedNameSyntaxID,
		Names:    StringArray{"distinguishedNameMatch"},
	},
	MatchingRule{
		OID:      GeneralizedTimeMatchRuleID,
		SyntaxID: GeneralizedTimeSyntaxID,
		Names:    StringArray{"generalizedTimeMatch"},
	},
	MatchingRule{
		OID:      GeneralizedTimeOrderingMatchRuleID,
		SyntaxID: GeneralizedTimeSyntaxID,
		Names:    StringArray{"generalizedTimeOrderingMatch"},
	},
	MatchingRule{
		OID:      IntegerFirstCompMatchRuleID,
		SyntaxID: IntegerSyntaxID,
		Names:    StringArray{"integerFirstComponentMatch"},
	},
	MatchingRule{
		OID:      IntegerMatchRuleID,
		SyntaxID: IntegerSyntaxID,
		Names:    StringArray{"integerMatch"},
	},
	MatchingRule{
		OID:      IntegerOrderingMatchRuleID,
		SyntaxID: IntegerSyntaxID,
		Names:    StringArray{"integerOrderingMatch"},
	},
	MatchingRule{
		OID:      KeywordMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"keywordMatch"},
	},
	MatchingRule{
		OID:      NumericStringMatchRuleID,
		SyntaxID: NumericStringSyntaxID,
		Names:    StringArray{"numericStringMatch"},
	},
	MatchingRule{
		OID:      NumericStringOrderingMatchRuleID,
		SyntaxID: NumericStringSyntaxID,
		Names:    StringArray{"numericStringOrderingMatch"},
	},
	MatchingRule{
		OID:      NumericStringSubstrMatchRuleID,
		SyntaxID: SubstringAssertionSyntaxID,
		Names:    StringArray{"numericStringSubstringsMatch"},
	},
	MatchingRule{
		OID:      ObjectIdentifierFirstCompMatchRuleID,
		SyntaxID: OIDSyntaxID,
		Names:    StringArray{"objectIdentifierFirstComponentMatch"},
	},
	MatchingRule{
		OID:      ObjectIdentifierMatchRuleID,
		SyntaxID: OIDSyntaxID,
		Names:    StringArray{"objectIdentifierMatch"},
	},
	MatchingRule{
		OID:      OctetStringMatchRuleID,
		SyntaxID: OctetStringSyntaxID,
		Names:    StringArray{"octetStringMatch"},
	},
	MatchingRule{
		OID:      OctetStringOrderingMatchRuleID,
		SyntaxID: OctetStringSyntaxID,
		Names:    StringArray{"octetStringOrderingMatch"},
	},
	MatchingRule{
		OID:      TelephoneNumberMatchRuleID,
		SyntaxID: TelephoneNumberSyntaxID,
		Names:    StringArray{"telephoneNumberMatch"},
	},
	MatchingRule{
		OID:      TelephoneNumberSubstrMatchRuleID,
		SyntaxID: SubstringAssertionSyntaxID,
		Names:    StringArray{"telephoneNumberSubstringsMatch"},
	},
	MatchingRule{
		OID:      UniqueMemberMatchRuleID,
		SyntaxID: NameAndOptionalUIDSyntaxID,
		Names:    StringArray{"uniqueMemberMatch"},
	},
	MatchingRule{
		OID:      WordMatchRuleID,
		SyntaxID: DirectoryStringSyntaxID,
		Names:    StringArray{"wordMatch"},
	},
}
