package models

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// HashIterations is the number of iterations (used by PBKDF2)
	hashIterations = 4096
	// HashKeyLength is the desired derived key length (used by PBKDF2)
	hashKeyLength = 32
	saltSize      = 16
)

// User model in the DB
type User struct {
	Id           int64
	Created      int64
	Username     string
	PasswordHash string
	PasswordSalt string
}

// CreateUser creates a User with the specified username and password
func CreateUser(username string, password string) User {
	user := User{
		Created:  time.Now().UnixNano(),
		Username: username,
	}
	user.SetPassword(password)
	return user
}

// ComparePassword compares the password with the user's hash and salt
func (user *User) ComparePassword(password string) (result bool, err error) {
	salt, err := base64.StdEncoding.DecodeString(user.PasswordSalt)
	if err != nil {
		return false, fmt.Errorf("DecodeString failed: %v", err)
	}

	passwordHash := pbkdf2.Key([]byte(password), salt, hashIterations, hashKeyLength, sha1.New)
	expecting := base64.StdEncoding.EncodeToString(passwordHash)
	actual := user.PasswordHash

	return expecting == actual, nil
}

// SetPassword sets the password hash and salt on a user
func (user *User) SetPassword(password string) error {
	salt, err := generateSalt()
	if err != nil {
		return fmt.Errorf("generateSalt failed: %v", err)
	}

	passwordHash := pbkdf2.Key([]byte(password), salt, hashIterations, hashKeyLength, sha1.New)
	user.PasswordHash = base64.StdEncoding.EncodeToString(passwordHash)
	user.PasswordSalt = base64.StdEncoding.EncodeToString(salt)
	return nil
}

func generateSalt() (result []byte, err error) {
	salt := make([]byte, saltSize)
	_, err = io.ReadFull(rand.Reader, salt)
	return salt, err
}
