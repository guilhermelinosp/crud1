package cryptography

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/guilhermelinosp/crud1/config/logs"
)

func EncryptPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		logs.Error("An error occurred while generating salt: %v", err)
		return "", err
	}

	hash := generateHash(password, salt)
	return fmt.Sprintf("%s.%s", salt, hash), nil
}

func VerifyPassword(password, hashedPassword string) (bool, error) {
	parts := splitHashedPassword(hashedPassword)
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hashed password format")
	}

	salt := parts[0]
	hash := parts[1]
	hashedPasswordAttempt := generateHash(password, salt)

	return hash == hashedPasswordAttempt, nil
}

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		logs.Error("An error occurred while generating salt: %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func generateHash(password, salt string) string {
	hashBytes := sha256.Sum256([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(hashBytes[:])
}

func splitHashedPassword(hashedPassword string) []string {
	return strings.SplitN(hashedPassword, ".", 2)
}
