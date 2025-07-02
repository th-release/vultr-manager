package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func DecodeBase64(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// XorEncrypt encrypts the input text using XOR with the provided key and returns the result as a hexadecimal string.
func XorEncrypt(text, key string) string {
	textBytes := []byte(text) // Convert text to UTF-8 bytes
	keyBytes := []byte(key)   // Convert key to UTF-8 bytes
	result := strings.Builder{}

	for i := 0; i < len(textBytes); i++ {
		// XOR operation with the key byte at the corresponding position
		c := textBytes[i] ^ keyBytes[i%len(keyBytes)]
		// Convert to 2-digit hexadecimal and append
		result.WriteString(hex.EncodeToString([]byte{c}))
	}
	return result.String()
}

// XorDecrypt decrypts the hexadecimal input text using XOR with the provided key and returns the original string.
func XorDecrypt(hexText, key string) (string, error) {
	keyBytes := []byte(key) // Convert key to UTF-8 bytes
	// Decode hexadecimal string to bytes
	bytes, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}

	// XOR operation to decrypt
	result := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		result[i] = bytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	return string(result), nil // Convert bytes back to UTF-8 string
}

func Sha512Hex(input string) string {
	hash := sha512.Sum512([]byte(input)) // SHA-512 해시 계산
	return fmt.Sprintf("%x", hash)       // 16진수 문자열로 변환 후 반환
}
