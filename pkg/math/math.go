package math


import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateHashLink(data string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	currentTime := time.Now().UnixNano()

	input := fmt.Sprintf("%x-%d-%s", salt, currentTime, data)

	hash := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hash[:]), nil
}