package main

import (
	"fmt"
	"strings"
)

func delMenu() {
	for {
		clearTerminal(true)
		fmt.Println("GoVault delete menu")
		fmt.Println("")
		fmt.Println("1 - Delet by SID")
		fmt.Println("2 - Delet by Search")
		fmt.Println("3 - Return")
		fmt.Println()

		menuCode := getStrInput("Select 1-3")
		clearTerminal(true)

		switch menuCode {
		case "1":
			delOpt1()
		case "2":
			delOpt2()
		case "3":
			delOpt3()
		}

	}
}

func delOpt1() {
	clearTerminal(true)
	//Get SID to del
	fmt.Println("Enter the SID to delet")
	delSID, _ := getIntInput("SID")
	foundPWData := dbGetDataBySID(delSID)

	if !doesSidExist(foundPWData) || delSID == 0 {
		fmt.Println("Ther is no SID", delSID)
		pressEnterToContinue()
		return
	}

	//Print found entry
	clearTerminal(true)
	formatPwData(foundPWData, "")

	if getDecision("Do you want to delete this password?") {
		dbDelete(delSID)
		printEventMessage("You have successfully deleted this password!")
		mainMenu()
	} else {
		return
	}

}

func delOpt2() {
	search := ""
	for {
		search = getStrInput("Search for passwords to delete")
		if strings.TrimSpace(search) != "" {
			break
		}
		clearTerminal(true)
	}
	delData := dbGetDataBySearch(search)

	if len(delData) == 0 {
		clearTerminal(true)
		printEventMessage("No passwords found for this search!")
		return
	}
	clearTerminal(true)
	formatTable(delData)
	fmt.Println()

	if getDecision("Would you like to delete all these passwords?") {

		for _, delPW := range delData {
			dbDelete(delPW.SID)
		}

		printEventMessage("All passwords have been deleted!")
		mainMenu()

	} else {
		return
	}

}

func delOpt3() {
	mainMenu()
}
