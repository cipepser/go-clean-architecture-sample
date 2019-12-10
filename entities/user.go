package entities

import (
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (u *User) PasswordVerify(pw string) bool {
	s := strings.Split(u.Password, "$")
	cost, err := strconv.Atoi(s[1])
	if err != nil {
		return false
	}
	hash := pbkdf2.Key([]byte(pw), []byte(s[2]), cost, sha256.Size, sha256.New)
	b64Hash := base64.StdEncoding.EncodeToString(hash)

	return b64Hash == s[3]
}
