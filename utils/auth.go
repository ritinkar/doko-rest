package utils

import (
	"log"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt : take a string and hash it
func HashAndSalt(plainPwd string) string {

	bytePwd := []byte(plainPwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {

	bytePwd := []byte(plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GenerateSalt : generates random salt
func GenerateSalt() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
