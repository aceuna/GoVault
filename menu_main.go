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
			mainOpt0()
		case "1":
			mainOpt1()
		case "2":
			mainOpt2()
		case "3":
			mainOpt3()
		case "4":
			mainOpt4()
		case "5":
			mainOpt5()
		case "6":
			mainOpt6()
		case "7":
			mainOpt7()
		}

	}

}

func mainOpt0() {

	clearTerminal(true)
	//Get SID
	fmt.Println("Enter the SID")
	searchSID, _ := getIntInput("SID")
	foundPWData := dbGetDataBySID(searchSID)

	if !doesSidExist(foundPWData) || searchSID == 0 {
		fmt.Println("Ther is no SID", searchSID)
		pressEnterToContinue()
		return
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
		pressEnterToContinue()
	}

}

func mainOpt1() {
	dbData := dbRead()
	dbData = dbData[1:]
	fmt.Println("These are your passwords")
	fmt.Println()
	formatTable(dbData)
	//dbGetALLData()
	pressEnterToContinue()
}

func mainOpt2() {
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

func mainOpt3() {
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
		printEventMessage("Your password has been saved!")
	}

}

func mainOpt4() {
	clearTerminal(true)
	//Get SID to del
	fmt.Println("Enter the SID to edit")
	delSID, NoInput := getIntInput("SID")

	if NoInput == "NoInput" {
		return
	}

	foundPWData := dbGetDataBySID(delSID)

	if !doesSidExist(foundPWData) || delSID == 0 {
		fmt.Println("Ther is no SID", delSID)
		pressEnterToContinue()
		return
	}

	//Print found entry
	clearTerminal(true)
	formatPwData(foundPWData, "")

	if getDecision("Do you want to edit this password?") {
		editMenu(foundPWData)
	}
}

func mainOpt5() {
	delMenu()
}

func mainOpt6() {
	login()
}
func mainOpt7() {
	exit()
}
