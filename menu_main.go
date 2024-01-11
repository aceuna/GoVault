package main

import (
	"fmt"
	"strings"
)

func mainMenu() {
	for {
		clearTerminal(true)
		printWelcome()

		fmt.Println("0 - Get Password")
		fmt.Println("1 - Show all Passwords")
		fmt.Println("2 - Search Password")
		fmt.Println("3 - Add Password")
		fmt.Println("4 - Edit Password")
		fmt.Println("5 - Delete Password")
		fmt.Println("6 - Logout")
		fmt.Println("7 - EXIT")
		fmt.Println("")

		menuCode := getStrInput("Select 0-7")
		clearTerminal(true)

		switch menuCode {
		case "0":
			opt0()
		case "1":
			opt1()
		case "2":
			opt2()
		case "3":
			opt3()
		case "4":
			opt4()
		case "5":
			opt5()
		case "6":
			opt6()
		case "7":
			opt7()
		}

	}

}

func opt0() {

	clearTerminal(true)
	//Get SID
	fmt.Println("Enter the SID")
	searchSID, _ := getIntInput("SID")
	foundPWData := dbGetDataBySID(searchSID)

	if !doesSidExist(foundPWData) || searchSID == 0 {
		fmt.Println("Ther is no SID", searchSID)
		pressEnterToContinue()
		mainMenu()
	}

	//Print found entry
	clearTerminal(true)
	formatPwData(foundPWData, "")

	if getDecision("Do you want to see this password?") {
		pw, check := checkPwWithHash()

		if check {
			clearTerminal(true)
			key := deriveKey(pw)
			encryptedData := foundPWData.Password
			// Decrypt
			decryptedPW, _ := decrypt(encryptedData, key)

			formatPwData(foundPWData, decryptedPW)

		} else {
			fmt.Println("Wrong Password!")
			pressEnterToContinue()
			mainMenu()
		}
	} else {
		mainMenu()
	}
	pressEnterToContinue()
}

func opt1() {
	dbData := dbRead()
	dbData = dbData[1:]
	fmt.Println("These are your passwords")
	fmt.Println()
	formatTable(dbData)
	//dbGetALLData()
	pressEnterToContinue()
}

func opt2() {
	search := ""
	for {
		search = getStrInput("Search")
		if strings.TrimSpace(search) != "" {
			break
		}
		clearTerminal(true)
	}
	formatTable(dbGetDataBySearch(search))
	pressEnterToContinue()
}

func opt3() {
	fmt.Println("Add a new PW")
	newUser, newPW, newURL, newNote := "", "", "", ""
	newUser = getStrInput("Username")
	newPW = getStrInput("Password")
	newURL = getStrInput("URL")
	newNote = getStrInput("NOTE")
	newSID := generateNewSID()
	fmt.Println("")

	if getDecision("Do you want to save this password") {
		for {
			pw, check := checkPwWithHash()
			if check {
				key := deriveKey(pw)
				newPW, _ = encrypt([]byte(newPW), key)
				break
			}
		}

		newData := []PWData{
			{SID: newSID, Username: newUser, Password: newPW, URL: newURL, Note: newNote},
		}
		dbAddData(newData)
	} else {
		mainMenu()
	}

}

func opt4() {
	clearTerminal(true)
	//Get SID to del
	fmt.Println("Enter the SID to edit")
	delSID, NoInput := getIntInput("SID")

	if NoInput == "NoInput" {
		mainMenu()
	}

	foundPWData := dbGetDataBySID(delSID)

	if !doesSidExist(foundPWData) || delSID == 0 {
		fmt.Println("Ther is no SID", delSID)
		pressEnterToContinue()
		mainMenu()
	}

	//Print found entry
	clearTerminal(true)
	formatPwData(foundPWData, "")

	if getDecision("Do you want to edit this password?") {
		editMenu(foundPWData)
	} else {
		mainMenu()
	}

}

func opt5() {
	delMenu()
}

func opt6() {
	login()
}
func opt7() {
	exit()
}
