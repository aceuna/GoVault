package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func createMD5Hash(data string) string {
	// Convert the string to bytes
	byte := []byte(data)

	// Create an MD5 hash
	hash := md5.Sum(byte)

	// Convert the hash to a hex string
	newHash := hex.EncodeToString(hash[:])

	return newHash
}

func deriveKey(password string) []byte {
	// Generate a key using PBKDF2
	return pbkdf2.Key([]byte(password), nil, 10000, 32, sha256.New)
}

func encrypt(data, key []byte) (string, error) {
	// Create a new cipher block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// Create a new GCM block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// Encrypt the data using the GCM block and the nonce
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(data string, key []byte) (string, error) {
	// Decode the data from base64
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	// Create a new cipher block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// Create a new GCM block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// Get the nonce size
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	// Get the nonce and the ciphertext from the data
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	// Return the decrypted data
	return string(plaintext), nil
}
