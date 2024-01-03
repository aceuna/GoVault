package main

import (
	"fmt"
	"strings"
	"time"
)

func login() {
	clearTerminal()

	//create login if none exists
	if !fileExists(filePath) {

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
			{SID: 0, Username: user, Password: createMD5Hash(pw), URL: "", Note: ""},
		}
		dbAppend(newUser)
		login()
		return
	}
	fmt.Println("Welcome to GoVault! by Steven™")
	for {

		loginUser := getStrInput("User")
		loginPw := getStrInput("Password")

		if strings.EqualFold(loginUser, dbGetDataBySID(0).Username) && createMD5Hash(loginPw) == dbGetDataBySID(0).Password {
			break
		} else {
			fmt.Println("You have entered the wrong user or password!!")
			pressEnter()
			login()
			return
		}

	}
	clearTerminal()
	//fmt.Println("Welcome to your GoVault " + dbGetDataBySID(0).Username + "!")
	loading()
	time.Sleep(1500 * time.Millisecond)
}
