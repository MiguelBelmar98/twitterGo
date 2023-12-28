package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(passwd string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), costo)
	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil
}
