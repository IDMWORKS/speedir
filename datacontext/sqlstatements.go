package datacontext

const (
	// Users table
	sqlCreateUsersTable = `
CREATE TABLE IF NOT EXISTS users
(
  id bigserial NOT NULL PRIMARY KEY,
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
	SqlSelectUserByUsername = `
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

	sqlInsertSyntaxRow = `
INSERT INTO syntaxes
(oid, description)
VALUES
($1, $2)`

	// MatchingRules table
	sqlCreateMatchingRulesTable = `
CREATE TABLE IF NOT EXISTS matching_rules
(
	oid text PRIMARY KEY,
	syntax_id text NOT NULL REFERENCES syntaxes (oid),
	names text[] NOT NULL
)
WITH (
	OIDS=FALSE
)`
	sqlSelectMatchingRuleCount = `
SELECT COUNT(oid) FROM matching_rules`
	sqlInsertMatchingRuleRow = `
INSERT INTO matching_rules
(oid, syntax_id, names)
VALUES
($1, $2, $3)`

	// AttributeTypes table
	sqlCreateAttributeTypesTable = `
CREATE TABLE IF NOT EXISTS attribute_types
(
	oid text PRIMARY KEY,
	super_id text REFERENCES attribute_types (oid),
	syntax_id text REFERENCES syntaxes (oid),
	names text[] NOT NULL,
	flags int NOT NULL,
	equality_match_id text REFERENCES matching_rules (oid),
	substring_match_id text REFERENCES matching_rules (oid),
	ordering_match_id text REFERENCES matching_rules (oid)
)
WITH (
	OIDS=FALSE
)`
	sqlSelectAttributeTypeCount = `
SELECT COUNT(oid) FROM attribute_types`
	sqlInsertAttributeTypeRow = `
INSERT INTO attribute_types
(oid, syntax_id, super_id, names, flags,
	equality_match_id, substring_match_id, ordering_match_id)
VALUES
($1, $2, $3, $4, $5,
	$6, $7, $8)`

	// ObjectClasses table
	sqlCreateObjectClassesTable = `
CREATE TABLE IF NOT EXISTS object_classes
(
	oid text PRIMARY KEY,
	super_id text REFERENCES object_classes (oid),
	names text[] NOT NULL,
	flags int NOT NULL,
	must_attribute_ids text[],
	may_attribute_ids text[]
)
WITH (
	OIDS=FALSE
)`
	sqlSelectObjectClassCount = `
SELECT COUNT(oid) FROM object_classes`
	sqlInsertObjectClassRow = `
INSERT INTO object_classes
(oid, super_id, names, flags,
	must_attribute_ids, may_attribute_ids)
VALUES
($1, $2, $3, $4,
	$5, $6)`
)
