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
	PW_MIN_LEN    = 5
)

type User struct {
	BaseModel `bson:",inline"`
	Active    bool
	Email     string
	Password  string
	Salt      string
}

func (user User) AuthPassword(pass string) bool {
	if !user.Active {
		return false
	}
	hashedPassword := string(pbkdf2.Key(
		[]byte(pass),
		[]byte(user.Salt),
		PW_HASH_ITER,
		PW_HASH_BYTES,
		sha512.New,
	))

	return hashedPassword == user.Password
}

func (user *User) CreatePassword(pass string) error {
	if len(pass) < PW_MIN_LEN {
		return fmt.Errorf("Password must have minimum length of %d", PW_MIN_LEN)
	}
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	user.Salt = string(salt)
	user.Password = string(pbkdf2.Key(
		[]byte(pass),
		salt,
		PW_HASH_ITER,
		PW_HASH_BYTES,
		sha512.New,
	))

	return nil
}

// FIXME: #16
func (user User) Validate() error {
	if user.Email == "" {
		return ErrNoEmpty("Email")
	}
	if user.Password == "" {
		return ErrNoEmpty("Password")
	}
	if user.Salt == "" {
		return ErrNoEmpty("Salt")
	}
	return nil
}

func (user User) String() string {
	return fmt.Sprintf("<User %s>", user.Email)
}

type UserInput struct {
	Email    *string
	Password *string
}
