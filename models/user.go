package models

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"

	"github.com/nwoolls/speedir/errors"
	"golang.org/x/crypto/pbkdf2"
)

const (
	//HashIterations is the number of iterations (used by PBKDF2)
	hashIterations = 4096
	//HashKeyLength is the desired derived key length (used by PBKDF2)
	hashKeyLength = 32
	saltSize      = 16
)

//User model in the DB
type User struct {
	Id           int64
	Created      int64
	Username     string
	PasswordHash string
	PasswordSalt string
}

//ComparePassword compares the password with the user's hash and salt
func (user *User) ComparePassword(password string) bool {
	salt, err := base64.StdEncoding.DecodeString(user.PasswordSalt)
	errors.CheckErr(err, "DecodeString failed")

	passwordHash := pbkdf2.Key([]byte(password), salt, hashIterations, hashKeyLength, sha1.New)
	expecting := base64.StdEncoding.EncodeToString(passwordHash)
	actual := user.PasswordHash

	return expecting == actual
}

//SetPassword sets the password hash and salt on a user
func (user *User) SetPassword(password string) {
	salt := generateSalt()
	passwordHash := pbkdf2.Key([]byte(password), salt, hashIterations, hashKeyLength, sha1.New)
	user.PasswordHash = base64.StdEncoding.EncodeToString(passwordHash)
	user.PasswordSalt = base64.StdEncoding.EncodeToString(salt)
}

func generateSalt() []byte {
	salt := make([]byte, saltSize)
	_, err := io.ReadFull(rand.Reader, salt)
	errors.CheckErr(err, "Random read failed")

	return salt
}
