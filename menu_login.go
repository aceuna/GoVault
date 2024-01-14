package main

import (
	"fmt"
	"strings"
	"time"
)

func loginMenu() {
	clearTerminal(true)

	//create login if none exists
	if !fileExists(jsonFilePath) {
		creatLogin(0)
		loginMenu()
		return
	}
	//normal login
	fmt.Println("Welcome to GoVault!")
	fmt.Println()
	for {

		loginUser := getStrInput("User")
		loginPw := getStrInput("Password")

		if strings.EqualFold(loginUser, dbGetDataBySID(0).Username) && createMD5Hash(loginPw) == dbGetDataBySID(0).Password {
			break
		} else {
			fmt.Println("You have entered the wrong user or password!!")
			pressEnterToContinue()
			loginMenu()
			return
		}

	}
	//Login
	loginMessage()
}

func creatLogin(SID int) {
	fmt.Println("Welcome to GoVault. Create your user!")
	user := getStrInput("User")
	var pw string
	//get right PW
	for {
		pw = getStrInput("Password")
		if pw == getStrInput("Repeat password") {
			break
		}
	}
	newUser := []PWData{
		{
			SID:      SID,
			Username: user,
			Password: createMD5Hash(pw),
			URL:      "https://govault.ch/",
			Note:     "System PW",
		},
	}
	dbAddData(newUser)
}

func loginMessage() {
	clearTerminal(true)
	printWelcome()
	printLogin()
	time.Sleep(1500 * time.Millisecond)
}
