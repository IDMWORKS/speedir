package models

// MatchingRule model in the DB
type MatchingRule struct {
	Name   string
	OID    string
	Syntax string
	Names  StringSlice
}

const (
	// OIDs
	// ObjectIdentifierMatchRuleID ... http://tools.ietf.org/html/rfc4517
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

	// names
	// ObjectIdentifierMatchRule ... http://tools.ietf.org/html/rfc4517
	ObjectIdentifierMatchRule          = "objectIdentifierMatch"
	DistinguishedNameMatchRule         = "distinguishedNameMatch"
	CaseIgnoreMatchRule                = "caseIgnoreMatch"
	CaseIgnoreOrderingMatchRule        = "caseIgnoreOrderingMatch"
	CaseIgnoreSubstrMatchRule          = "caseIgnoreSubstringsMatch"
	CaseExactMatchRule                 = "caseExactMatch"
	CaseExactOrderingMatchRule         = "caseExactOrderingMatch"
	CaseExactSubstrMatchRule           = "caseExactSubstringsMatch"
	NumericStringMatchRule             = "numericStringMatch"
	NumericStringOrderingMatchRule     = "numericStringOrderingMatch"
	NumericStringSubstrMatchRule       = "numericStringSubstringsMatch"
	CaseIgnoreListMatchRule            = "caseIgnoreListMatch"
	CaseIgnoreListSubstrMatchRule      = "caseIgnoreListSubstringsMatch"
	BooleanMatchRule                   = "booleanMatch"
	IntegerMatchRule                   = "integerMatch"
	IntegerOrderingMatchRule           = "integerOrderingMatch"
	BitStringMatchRule                 = "bitStringMatch"
	OctetStringMatchRule               = "octetStringMatch"
	OctetStringOrderingMatchRule       = "octetStringOrderingMatch"
	TelephoneNumberMatchRule           = "telephoneNumberMatch"
	TelephoneNumberSubstrMatchRule     = "telephoneNumberSubstringsMatch"
	UniqueMemberMatchRule              = "uniqueMemberMatch"
	GeneralizedTimeMatchRule           = "generalizedTimeMatch"
	GeneralizedTimeOrderingMatchRule   = "generalizedTimeOrderingMatch"
	IntegerFirstCompMatchRule          = "integerFirstComponentMatch"
	ObjectIdentifierFirstCompMatchRule = "objectIdentifierFirstComponentMatch"
	DirectoryStringFirstCompMatchRule  = "directoryStringFirstComponentMatch"
	KeywordMatchRule                   = "keywordMatch"
	WordMatchRule                      = "wordMatch"
	CaseExactIA5MatchRule              = "caseExactIA5Match"
	CaseIgnoreIA5MatchRule             = "caseIgnoreIA5Match"
	CaseIgnoreIA5SubstrMatchRule       = "caseIgnoreIA5SubstringsMatch"
)

// LDAPv3MatchingRules represents the standard LDAPv3 Matching Rules
// http://www.ietf.org/rfc/rfc4517.txt
var LDAPv3MatchingRules = [...]MatchingRule{
	MatchingRule{
		OID:    BitStringMatchRuleID,
		Syntax: BitStringSyntaxID,
		Name:   BitStringMatchRule,
	},
	MatchingRule{
		OID:    BooleanMatchRuleID,
		Syntax: BooleanSyntaxID,
		Name:   BooleanMatchRule,
	},
	MatchingRule{
		OID:    CaseExactIA5MatchRuleID,
		Syntax: IA5StringSyntaxID,
		Name:   CaseExactIA5MatchRule,
	},
	MatchingRule{
		OID:    CaseExactMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   CaseExactMatchRule,
	},
	MatchingRule{
		OID:    CaseExactOrderingMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   CaseExactOrderingMatchRule,
	},
	MatchingRule{
		OID:    CaseExactSubstrMatchRuleID,
		Syntax: SubstringAssertionSyntaxID,
		Name:   CaseExactSubstrMatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreIA5MatchruleID,
		Syntax: IA5StringSyntaxID,
		Name:   CaseIgnoreIA5MatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreIA5SubstrMatchRuleID,
		Syntax: SubstringAssertionSyntaxID,
		Name:   CaseIgnoreIA5SubstrMatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreListMatchRuleID,
		Syntax: PostalAddressSyntaxID,
		Name:   CaseIgnoreListMatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreListSubstrMatchRuleID,
		Syntax: SubstringAssertionSyntaxID,
		Name:   CaseIgnoreListSubstrMatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   CaseIgnoreMatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreOrderingMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   CaseIgnoreOrderingMatchRule,
	},
	MatchingRule{
		OID:    CaseIgnoreSubstrMatchRuleID,
		Syntax: SubstringAssertionSyntaxID,
		Name:   CaseIgnoreSubstrMatchRule,
	},
	MatchingRule{
		OID:    DirectoryStringFirstCompMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   DirectoryStringFirstCompMatchRule,
	},
	MatchingRule{
		OID:    DistinguishedNameMatchRuleID,
		Syntax: DistinguishedNameSyntaxID,
		Name:   DistinguishedNameMatchRule,
	},
	MatchingRule{
		OID:    GeneralizedTimeMatchRuleID,
		Syntax: GeneralizedTimeSyntaxID,
		Name:   GeneralizedTimeMatchRule,
	},
	MatchingRule{
		OID:    GeneralizedTimeOrderingMatchRuleID,
		Syntax: GeneralizedTimeSyntaxID,
		Name:   GeneralizedTimeOrderingMatchRule,
	},
	MatchingRule{
		OID:    IntegerFirstCompMatchRuleID,
		Syntax: IntegerSyntaxID,
		Name:   IntegerFirstCompMatchRule,
	},
	MatchingRule{
		OID:    IntegerMatchRuleID,
		Syntax: IntegerSyntaxID,
		Name:   IntegerMatchRule,
	},
	MatchingRule{
		OID:    IntegerOrderingMatchRuleID,
		Syntax: IntegerSyntaxID,
		Name:   IntegerOrderingMatchRule,
	},
	MatchingRule{
		OID:    KeywordMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   KeywordMatchRule,
	},
	MatchingRule{
		OID:    NumericStringMatchRuleID,
		Syntax: NumericStringSyntaxID,
		Name:   NumericStringMatchRule,
	},
	MatchingRule{
		OID:    NumericStringOrderingMatchRuleID,
		Syntax: NumericStringSyntaxID,
		Name:   NumericStringOrderingMatchRule,
	},
	MatchingRule{
		OID:    NumericStringSubstrMatchRuleID,
		Syntax: SubstringAssertionSyntaxID,
		Name:   NumericStringSubstrMatchRule,
	},
	MatchingRule{
		OID:    ObjectIdentifierFirstCompMatchRuleID,
		Syntax: OIDSyntaxID,
		Name:   ObjectIdentifierFirstCompMatchRule,
	},
	MatchingRule{
		OID:    ObjectIdentifierMatchRuleID,
		Syntax: OIDSyntaxID,
		Name:   ObjectIdentifierMatchRule,
	},
	MatchingRule{
		OID:    OctetStringMatchRuleID,
		Syntax: OctetStringSyntaxID,
		Name:   OctetStringMatchRule,
	},
	MatchingRule{
		OID:    OctetStringOrderingMatchRuleID,
		Syntax: OctetStringSyntaxID,
		Name:   OctetStringOrderingMatchRule,
	},
	MatchingRule{
		OID:    TelephoneNumberMatchRuleID,
		Syntax: TelephoneNumberSyntaxID,
		Name:   TelephoneNumberMatchRule,
	},
	MatchingRule{
		OID:    TelephoneNumberSubstrMatchRuleID,
		Syntax: SubstringAssertionSyntaxID,
		Name:   TelephoneNumberSubstrMatchRule,
	},
	MatchingRule{
		OID:    UniqueMemberMatchRuleID,
		Syntax: NameAndOptionalUIDSyntaxID,
		Name:   UniqueMemberMatchRule,
	},
	MatchingRule{
		OID:    WordMatchRuleID,
		Syntax: DirectoryStringSyntaxID,
		Name:   WordMatchRule,
	},
}
