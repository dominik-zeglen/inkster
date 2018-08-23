package core

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
	PW_HASH_ITER  = 4096
)

type User struct {
	BaseModel
	Active   bool
	Email    string
	password string
	salt     string
}

func (user User) AuthPassword(pass string) bool {
	if !user.Active {
		return false
	}
	hashedPassword := string(pbkdf2.Key(
		[]byte(pass),
		[]byte(user.salt),
		PW_HASH_ITER,
		PW_HASH_BYTES,
		sha512.New,
	))

	return hashedPassword == user.password
}

func (user *User) CreatePassword(pass string) *User {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	user.salt = string(salt)
	user.password = string(pbkdf2.Key(
		[]byte(pass),
		salt,
		PW_HASH_ITER,
		PW_HASH_BYTES,
		sha512.New,
	))

	return user
}

// FIXME: #16
func (user User) Validate() error {
	return nil
}

func (user User) String() string {
	return fmt.Sprintf("<User %s>", user.Email)
}

type UserInput struct {
	Email    string
	Password string
}
