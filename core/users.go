package core

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"

	"github.com/dchest/uniuri"
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
	Active    bool   `json:"active"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Salt      string `json:"salt" validate:"required"`
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

func (user *User) CreateRandomPassword() (string, error) {
	pwd := uniuri.NewLen(8)
	return string(pwd), user.CreatePassword(string(pwd))
}

func (user *User) CreatePassword(pass string) error {
	if len(pass) < PW_MIN_LEN {
		return fmt.Errorf("Password must have minimum length of %d", PW_MIN_LEN)
	}
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return err
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
func (user User) Validate() []ValidationError {
	return ValidateModel(user)
}

func (user User) String() string {
	return fmt.Sprintf("<User %s>", user.Email)
}

type UserInput struct {
	Active   *bool
	Email    *string
	Password *string
}
