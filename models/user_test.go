package models

import (
	"crypto/sha1"
	"encoding/base64"
	"testing"

	"golang.org/x/crypto/pbkdf2"
	"log"
)

func TestCreateUser(t *testing.T) {
	expected := "password"
	user := CreateUser("username", expected)
	if !comparePassword(user, expected) {
		t.Error("For", user, "comparePassword returned false")
	}
}

func TestSetPassword(t *testing.T) {
	exptected := "password2"
	user := CreateUser("username", "password")
	user.SetPassword(exptected)
	if !comparePassword(user, exptected) {
		t.Error("For", user, "comparePassword returned false")
	}
}

func TestComparePassword(t *testing.T) {
	expected := "password"
	user := CreateUser("username", expected)
	if match, _ := user.ComparePassword(expected); comparePassword(user, expected) != match {
		t.Error("For", user, "comparePassword != user.ComparePassword")
	}
}

func comparePassword(user User, password string) bool {
	salt, err := base64.StdEncoding.DecodeString(user.PasswordSalt)
	if err != nil {
		log.Fatalln("DecodeString failed", err)
	}

	passwordHash := pbkdf2.Key([]byte(password), salt, hashIterations, hashKeyLength, sha1.New)
	exptected := base64.StdEncoding.EncodeToString(passwordHash)

	return exptected == user.PasswordHash
}
