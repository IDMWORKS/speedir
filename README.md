# speedir
The speedir project is an exercise in learning both the Go programming language and the LDAPv3 protocol by implementing an LDAPv3 compatible server.

## Standards
* Lightweight Directory Access Protocol (LDAPv3): [RFC 4511](http://tools.ietf.org/html/rfc4511)
* LDAP Directory Information Models: [RFC 4512](http://tools.ietf.org/html/rfc4512)
* LDAP Syntaxes and Matching Rules: [RFC 4517](http://tools.ietf.org/html/rfc4517)
* LDAP Schema for User Applications: [RFC 4519](http://tools.ietf.org/html/rfc4519)
* Abstract Syntax Notation One (ASN.1 BER): [X.690](http://www.itu.int/ITU-T/studygroups/com17/languages/X.690-0207.pdf)
* Transport Layer Security (TLS): [RFC 5246](http://tools.ietf.org/html/rfc5246), [RFC 6176](http://tools.ietf.org/html/rfc6176)
* Password-Based Key Derivation Function 2 (PBKDF2, [FIPS 140-2](http://csrc.nist.gov/groups/STM/cmvp/documents/140-1/140val-all.htm) compliant): [RFC 2898](https://tools.ietf.org/html/rfc2898)

## Goals
* Simple codebase
* Fast
* Scalable

## Stack
* Go
* Postgres
* Docker (planned)
