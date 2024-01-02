package main

import (
	"fmt"
	"strings"
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
		return
	}
	fmt.Println("Welcome to GoVault!")
	for {

		loginUser := getStrInput("User")
		loginPw := getStrInput("Password")

		if strings.EqualFold(loginUser, dbGetDataBySID(0).Username) && createMD5Hash(loginPw) == dbGetDataBySID(0).Password {
			break
		} else {
			fmt.Println("You have entered the wrong user or password!!")
		}

	}

	fmt.Println("Welcome to your GoVault " + dbGetDataBySID(0).Username + "!")

}
