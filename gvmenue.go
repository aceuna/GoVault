package main

import (
	"fmt"
	"os"
	"strings"
)

func mainMenu() {
	for {

		clearTerminal()
		fmt.Println("GoVault main menu\n")
		fmt.Println("0 - Get Password")
		fmt.Println("1 - Show all Passwords")
		fmt.Println("2 - Search Password")
		fmt.Println("3 - Add Password")
		fmt.Println("4 - Edit Password")
		fmt.Println("5 - Delete Password")
		fmt.Println("6 - Logout")
		fmt.Println("7 - EXIT")

		menuCode := getStrInput("Select 1-6")
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
		fmt.Println("GoVault delete menu\n")
		fmt.Println("1 - Delet by SID")
		fmt.Println("2 - Delet by Search")
		fmt.Println("3 - Return")

		menuCode := getStrInput("Select 1-3")
		clearTerminal()
		switch menuCode {
		case "1":
			clearTerminal()
			//Get SID to del
			fmt.Println("Enter the SID to delet")
			delSID := getIntInput("SID")
			foundPWData := dbGetDataBySID(delSID)

			if !doesSidExist(foundPWData) {
				fmt.Println("Ther is no SID", delSID)
				getStrInput("")
				delMenu()
			}

			//Print found entry
			fmt.Println("SID, User, URL, Note")
			formatPWData(foundPWData)
			fmt.Println("\n")
			decision := getDecision("Do you want to delete this password? (y/n)")
			switch decision {
			case "y":
				dbDelete(delSID)
				//fmt.Println("The password has been deleted!")
			case "n":
				delMenu()
			}
			getStrInput("")
		case "2":

		case "3":
			mainMenu()
		}
	}
}

func opt0() {

	clearTerminal()
	//Get SID
	fmt.Println("Enter the SID")
	searchSID := getIntInput("SID")
	foundPWData := dbGetDataBySID(searchSID)

	if !doesSidExist(foundPWData) || searchSID == 0 {
		fmt.Println("Ther is no SID", searchSID)
		getStrInput("")
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
			getStrInput("")
			mainMenu()
		}

	case "n":
		main()
	}
	getStrInput("")
}

func opt1() {
	dbData := dbRead()
	dbData = dbData[1:]
	printTable(dbData)
	//dbGetALLData()
	getStrInput("")
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
	getStrInput("")
}

func opt3() {
	fmt.Println("Add a new PW")
	newUser, newPW, newURL, newNote := "", "", "", ""
	newUser = getStrInput("Username")
	newPW = getStrInput("Password")
	newURL = getStrInput("URL")
	newNote = getStrInput("NOTE")
	newSID := generateNewSID()

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

}

func opt4() {

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
