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
)
