package main

import (
	"fmt"
	"strconv"
)

var (
	newUser string
	newPW   string
	newURL  string
	newNote string
)

func editMenu(data PWData) {

	clearTerminal(true)

	fmt.Println("What do you want to edit?")
	fmt.Println()
	fmt.Println("SID:", strconv.Itoa(data.SID))
	fmt.Println()
	fmt.Println("1 - Username:", data.Username)
	fmt.Println("2 - Password:", "************")
	fmt.Println("3 - URL:", data.URL)
	fmt.Println("4 - Note:", data.Note)
	fmt.Println()

	editCode, _ := getIntInput("Select 1-4")
	fmt.Println()

	switch editCode {
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
			var pw string
			for {
				pw = getStrInput("New password")
				if pw == getStrInput("Repeat new password") {
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
	default:
		editMenu(data)
		return
	}

	replaceData := PWData{
		SID:      data.SID,
		Username: newUser,
		Password: newPW,
		URL:      newURL,
		Note:     newNote,
	}

	dbReplaceData(replaceData)
	printEventMessage("You have successfully edited this password!")
}
