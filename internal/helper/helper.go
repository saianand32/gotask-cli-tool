package helper

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateCryptoID creates a random 32-character hexadecimal string.
// It generates 16 bytes of random data and encodes it in hexadecimal format.
func GenerateCryptoID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GetMapValues extracts the values from a map and returns them as a slice.
// It works with any map type where K is a comparable type and V is any type.
func GetMapValues[K comparable, V any](m map[K]V) []V {
	var values []V
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
