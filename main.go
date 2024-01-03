package main

import (
	"os"
)

func main() {
	login()
	for {
		mainMenu()
	}

	//cryptostuff()

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
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
