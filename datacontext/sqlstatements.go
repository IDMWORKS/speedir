package datacontext

const (
	// Users table
	sqlCreateUsersTable = `
CREATE TABLE IF NOT EXISTS users
(
  id bigserial PRIMARY KEY,
  created bigint,
  username text,
  passwordhash text,
  passwordsalt text
)
WITH (
  OIDS=FALSE
);`
	sqlInsertUserRow = `
INSERT INTO users
(created, username, passwordhash, passwordsalt)
VALUES
($1, $2, $3, $4)`
	sqlSelectUserCountByUsername = `
SELECT COUNT(id) FROM users WHERE username = $1`
	sqlSelectUserByUsername = `
SELECT id, created, username, passwordhash, passwordsalt FROM users WHERE username = $1`

	// Syntaxes table
	sqlCreateSyntaxesTable = `
CREATE TABLE IF NOT EXISTS syntaxes
(
  oid text PRIMARY KEY,
  description text
)
WITH (
  OIDS=FALSE
);`
	sqlSelectSyntaxCount = `
SELECT COUNT(oid) FROM syntaxes`
	sqlSelectAllSyntaxes = `
SELECT oid, description FROM syntaxes`
	sqlInsertSyntaxRow = `
INSERT INTO syntaxes
(oid, description)
VALUES
($1, $2)`

	// MatchingRules table
	sqlCreateMatchingRulesTable = `
CREATE TABLE IF NOT EXISTS matching_rules
(
	name text PRIMARY KEY,
	oid text NOT NULL,
	syntax text NOT NULL REFERENCES syntaxes (oid),
	names text[]
)
WITH (
	OIDS=FALSE
)`
	sqlSelectMatchingRuleCount = `
SELECT COUNT(oid) FROM matching_rules`
	sqlInsertMatchingRuleRow = `
INSERT INTO matching_rules
(name, oid, syntax, names)
VALUES
($1, $2, $3, $4)`

	// AttributeTypes table
	sqlCreateAttributeTypesTable = `
CREATE TABLE IF NOT EXISTS attribute_types
(
	name text PRIMARY KEY,
	oid text NOT NULL,
	super text REFERENCES attribute_types (name),
	syntax text REFERENCES syntaxes (oid),
	names text[],
	flags int NOT NULL,
	usage int NOT NULL,
	equality_match text REFERENCES matching_rules (name),
	substring_match text REFERENCES matching_rules (name),
	ordering_match text REFERENCES matching_rules (name)
)
WITH (
	OIDS=FALSE
)`
	sqlSelectAttributeTypeCount = `
SELECT COUNT(oid) FROM attribute_types`
	sqlInsertAttributeTypeRow = `
INSERT INTO attribute_types
(name, oid, syntax, super, names, flags, usage,
	equality_match, substring_match, ordering_match)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// ObjectClasses table
	sqlCreateObjectClassesTable = `
CREATE TABLE IF NOT EXISTS object_classes
(
	name text PRIMARY KEY,
	oid text NOT NULL,
	super text REFERENCES object_classes (name),
	names text[],
	flags int NOT NULL,
	must_attributes text[],
	may_attributes text[]
)
WITH (
	OIDS=FALSE
)`
	sqlSelectObjectClassCount = `
SELECT COUNT(oid) FROM object_classes`
	sqlSelectAllObjectClasses = `
SELECT name
	, oid
	, super
	, array_to_json(names)
	, flags
	, array_to_json(must_attributes)
	, array_to_json(may_attributes)
FROM object_classes`
	sqlInsertObjectClassRow = `
INSERT INTO object_classes
(name, oid, super, names, flags,
	must_attributes, may_attributes)
VALUES
($1, $2, $3, $4, $5, $6, $7)`
)
