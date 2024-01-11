package main

import (
	"fmt"
	"strconv"
)

func editMenu(data PWData) {

	clearTerminal(true)

	fmt.Println("What do you want to edit?")
	fmt.Println()
	fmt.Println("SID:", strconv.Itoa(data.SID))
	fmt.Println("1 - Username:", data.Username)
	fmt.Println("2 - Password:", "************")
	fmt.Println("3 - URL:", data.URL)
	fmt.Println("4 - Note:", data.Note)
	fmt.Println()

	menuCode, _ := getIntInput("Select 1-4")
	var newUser, newPW, newURL, newNote string
	fmt.Println()
	switch menuCode {
	case 1:
		newUser = getStrInput("New Username")
		newPW = data.Password
		newURL = data.URL
		newNote = data.Note
	case 2:
		newUser = data.Username
		newURL = data.URL
		newNote = data.Note

		keyPw, check := checkPwWithHash()

		if check {
			pw := ""
			//get right PW
			for {
				pw = getStrInput("New password")
				pw2 := getStrInput("Repeat new password")
				if pw == pw2 {
					break
				}
			}

			key := deriveKey(keyPw)
			newPW, _ = encrypt([]byte(pw), key)

		} else {
			fmt.Println("Wrong Password!")
			pressEnterToContinue()
			editMenu(data)
		}

	case 3:
		newUser = data.Username
		newPW = data.Password
		newURL = getStrInput("New URL")
		newNote = data.Note
	case 4:
		newUser = data.Username
		newPW = data.Password
		newURL = data.URL
		newNote = getStrInput("New Note")
	}
	replaceData := PWData{
		SID: data.SID, Username: newUser, Password: newPW, URL: newURL, Note: newNote,
	}
	dbReplaceData(replaceData)
	printEventMessage("You have successfully edited this password!")
}
