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
	hashStr := hex.EncodeToString(hash[:])

	return hashStr
}

func deriveKey(password string) []byte {
	return pbkdf2.Key([]byte(password), nil, 10000, 32, sha256.New)
}

func encrypt(data, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(data string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

/*
func cryptostuff() {

	des := getStrInput("en=1,de=2,md5=3")
	password := getStrInput("PW")

	data := getStrInput("data")

	// Derive a 32-byte key from the password
	key := deriveKey(password)

	switch des {
	case "1":

		// Encrypt
		encryptedData, err := encrypt([]byte(data), key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}
		fmt.Println("Encrypted:", encryptedData)
	case "2":
		encryptedData := data
		// Decrypt
		decryptedData, err := decrypt(encryptedData, key)
		if err != nil {
			fmt.Println("Decryption error:", err)
			return
		}
		fmt.Println("Decrypted:", decryptedData)
	case "3":
		fmt.Println(data + " is " + createMD5Hash(data))

	}
}
*/
