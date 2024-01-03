package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainMenu() {
	for {

		clearTerminal()
		fmt.Println("Welcome to your GoVault " + dbGetDataBySID(0).Username + " !\n")

		//fmt.Println("GoVault main menu\n")
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
		clearTerminal()
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

func delMenu() {
	for {
		clearTerminal()
		fmt.Println("GoVault delete menu")
		fmt.Println("")
		fmt.Println("1 - Delet by SID")
		fmt.Println("2 - Delet by Search")
		fmt.Println("3 - Return")
		fmt.Println("")

		menuCode := getStrInput("Select 1-3")
		clearTerminal()
		switch menuCode {
		case "1":
			clearTerminal()
			//Get SID to del
			fmt.Println("Enter the SID to delet")
			delSID, _ := getIntInput("SID")
			foundPWData := dbGetDataBySID(delSID)

			if !doesSidExist(foundPWData) || delSID == 0 {
				fmt.Println("Ther is no SID", delSID)
				pressEnter()
				delMenu()
			}

			//Print found entry
			clearTerminal()
			formatPWDataWithPW(foundPWData, "*********")
			decision := getDecision("Do you want to delete this password? (y/n)")
			switch decision {
			case "y":
				dbDelete(delSID)
				//fmt.Println("The password has been deleted!")
			case "n":
				delMenu()
			}
			pressEnter()
		case "2":

		case "3":
			mainMenu()
		}
	}
}

func editMenu(data PWData) {

	clearTerminal()

	fmt.Println("Waht do you want to edit?")
	fmt.Println("")
	strNum := strconv.Itoa(data.SID)

	fmt.Println("SID:", strNum)
	fmt.Println("1 - Username:", data.Username)
	fmt.Println("2 - Password:", "***********")
	fmt.Println("3 - URL:", data.URL)
	fmt.Println("4 - Note:", data.Note)
	fmt.Println("")

	menuCode, _ := getIntInput("Select 1-4")
	var newUser, newPW, newURL, newNote string

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
				pw = getStrInput("Password")
				pw2 := getStrInput("Repeat password")
				if pw == pw2 {
					break
				}
			}

			key := deriveKey(keyPw)
			newPW, _ = encrypt([]byte(pw), key)

		} else {
			fmt.Println("Wrong Password!")
			pressEnter()
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
	pressEnter()
}

func opt0() {

	clearTerminal()
	//Get SID
	fmt.Println("Enter the SID")
	searchSID, _ := getIntInput("SID")
	foundPWData := dbGetDataBySID(searchSID)

	if !doesSidExist(foundPWData) || searchSID == 0 {
		fmt.Println("Ther is no SID", searchSID)
		pressEnter()
		mainMenu()
	}

	//Print found entry
	clearTerminal()
	formatPWDataWithPW(foundPWData, "***********")
	decision := getDecision("Do you want to see this password? (y/n)")
	switch decision {
	case "y":
		pw, check := checkPwWithHash()

		if check {
			clearTerminal()
			key := deriveKey(pw)
			encryptedData := foundPWData.Password
			// Decrypt
			decryptedPW, _ := decrypt(encryptedData, key)

			formatPWDataWithPW(foundPWData, decryptedPW)

		} else {
			fmt.Println("Wrong Password!")
			pressEnter()
			mainMenu()
		}

	case "n":
		mainMenu()
	}
	pressEnter()
}

func opt1() {
	dbData := dbRead()
	dbData = dbData[1:]
	printTable(dbData)
	//dbGetALLData()
	pressEnter()
}

func opt2() {
	search := ""
	for {
		search = getStrInput("Search")
		if strings.TrimSpace(search) != "" {
			break
		}
		clearTerminal()
	}
	printTable(dbGetSearch(search))
	pressEnter()
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

	decision := getDecision("Do you want to save this password (y/n)")

	switch decision {
	case "y":
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
		dbAppend(newData)
	case "n":
		mainMenu()
	}

}

func opt4() {
	clearTerminal()
	//Get SID to del
	fmt.Println("Enter the SID to edit")
	delSID, NoInput := getIntInput("SID")

	if NoInput == "NoInput" {
		mainMenu()
	}

	foundPWData := dbGetDataBySID(delSID)

	if !doesSidExist(foundPWData) || delSID == 0 {
		fmt.Println("Ther is no SID", delSID)
		pressEnter()
		mainMenu()
	}

	//Print found entry
	clearTerminal()
	formatPWDataWithPW(foundPWData, "***************")

	decision := getDecision("Do you want to edit this password? (y/n)")
	switch decision {
	case "y":
		editMenu(foundPWData)
	case "n":
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
	clearTerminal()
	os.Exit(1)
}
