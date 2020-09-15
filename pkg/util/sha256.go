package util

import (
	"crypto/sha256"
	"encoding/hex"
)

//EncodeToSHA256 encode the msg to sha256
func EncodeToSHA256(message string) string {
	bytes2 := sha256.Sum256([]byte(message))
	return hex.EncodeToString(bytes2[:])
}
