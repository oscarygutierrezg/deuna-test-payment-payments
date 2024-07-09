package auth

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base32"
	"fmt"
	"io"
	"os"
	"strings"
)

func EncryptPassword(password string) (string, string) {

	passwordByt := []byte(password)
	salt := generatePwdSalt(passwordByt)

	strSalt := strings.ToLower(base32.HexEncoding.EncodeToString(salt))
	strHash := applyHash(strSalt, password)

	return strSalt, strHash
}

func VerifyPassword(password string, strSalt string, strHash string) bool {
	strNewHash := applyHash(strSalt, password)
	return strNewHash == strHash
}

// private functions
func generatePwdSalt(secret []byte) []byte {

	const saltSize = 16

	buf := make([]byte, saltSize, saltSize+md5.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := md5.New()
	hash.Write(buf)
	hash.Write(secret)

	return hash.Sum(buf)
}

func applyHash(salt, secret string) string {
	return getSHA512(salt + secret)
}

func getSHA512(strVal string) string {
	h := sha512.New()
	h.Write([]byte(strVal))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
