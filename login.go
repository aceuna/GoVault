package main

import (
	"fmt"
	"strings"
	"time"
)

func login() {
	clearTerminal(true)

	//create login if none exists
	if !fileExists(jsonFilePath) {
		creatLogin(0)
		login()
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
			login()
			return
		}

	}
	//Login
	loginMessage()
}

func creatLogin(SID int) {
	fmt.Println("Welcome to GoVault. Create your user!")
	user := getStrInput("User")
	pw := ""
	//get right PW
	for {
		pw = getStrInput("Password")
		pw2 := getStrInput("Repeat password")
		if pw == pw2 {
			break
		}
	}
	newUser := []PWData{
		{SID: SID, Username: user, Password: createMD5Hash(pw), URL: "https://govault.ch/", Note: "System PW"},
	}
	dbAddData(newUser)
}

func loginMessage() {
	clearTerminal(true)
	printWelcome()
	printLogin()
	time.Sleep(1500 * time.Millisecond)
}
