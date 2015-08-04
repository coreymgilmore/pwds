/*
Package pwds is used to create and validate bcrypt passwords.
This is just a wrapper aroung golang.org/x/crypto/bcrypt that makes it easier to use.
*/

package pwds

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const BCRYPT_COST = 12

var ErrInvalidPassword = errors.New("invalidPassword")

//*********************************************************************************************************************************
//FUNCS

//CREATE A PASSWORD
//in: input (plain text password to hash)
//out: hashed password as a string
func Create(input string) string {
	raw := 		[]byte(input)
	hash, _ := 	bcrypt.GenerateFromPassword(raw, BCRYPT_COST)

	return string(hash)
}

//VALIDATE A PASSWORD
//in: rawPassword (plain text password), hash (hashed password from db)
//out: boolean, error if password is not valid
func Verify(rawPassword string, hashedPassword string) (bool, error) {
	rawPasswordByte := 	[]byte(rawPassword)
	hashPasswordByte := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashPasswordByte, rawPasswordByte)
	if err != nil {
		return false, ErrInvalidPassword
	}

	//password verified
	return true, nil
}
