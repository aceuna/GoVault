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
		fmt.Println("")

		menuCode := getStrInput("Select 1-3")
		clearTerminal(true)
		switch menuCode {
		case "1":
			clearTerminal(true)
			//Get SID to del
			fmt.Println("Enter the SID to delet")
			delSID, _ := getIntInput("SID")
			foundPWData := dbGetDataBySID(delSID)

			if !doesSidExist(foundPWData) || delSID == 0 {
				fmt.Println("Ther is no SID", delSID)
				pressEnterToContinue()
				delMenu()
			}

			//Print found entry
			clearTerminal(true)
			formatPwData(foundPWData, "")

			if getDecision("Do you want to delete this password?") {
				dbDelete(delSID)
				//fmt.Println("The password has been deleted!")
			} else {
				delMenu()
			}
			printEventMessage("You have successfully deleted this password!")
			mainMenu()
		case "2":
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
				delMenu()
			}
			clearTerminal(true)
			formatTable(delData)
			fmt.Println()

			if getDecision("Would you like to delete all these passwords?") {

				for _, delPW := range delData {
					dbDelete(delPW.SID)
				}

				printEventMessage("All passwords have been deleted!")

			} else {
				delMenu()
			}

			pressEnterToContinue()
			mainMenu()
		case "3":
			mainMenu()
		}
	}
}
