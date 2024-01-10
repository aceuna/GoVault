package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Get String Input
func getStrInput(inputMessage string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(inputMessage + ": ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal()
	}

	userInput = strings.TrimSpace(userInput)
	return userInput
}

// Get Int Input
func getIntInput(inputMessage string) (int, string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(inputMessage + ": ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}
		userInput = strings.TrimSpace(userInput)
		//filter no Input
		if userInput == "" {
			return 0, "NoInput"
		}
		if isInteger(userInput) {
			intInput, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				log.Fatal()
			}
			return intInput, ""
		}
	}
}

func formatPwData(data PWData, pw string) {
	if pw == "" {
		pw = "************"
	}

	fmt.Println("SID:", strconv.Itoa(data.SID))
	fmt.Println("Username:", data.Username)
	fmt.Println("Password:", pw)
	fmt.Println("URL:", data.URL)
	fmt.Println("Note:", data.Note)
	fmt.Println()
}

func newDecision(question string) string {
	var decision string
	for {
		decision = getStrInput(question + " (y/n)")
		decision = strings.ToLower(decision)
		if decision == "y" || decision == "n" {
			break
		}
	}
	return decision
}

func formatTable(PWData []PWData) {
	// Print table header
	fmt.Printf("%-5s %-25s %-35s %-40s\n", "SID", "USERNAME", "URL", "NOTE")
	// Print table rows
	for _, row := range PWData {

		var printUsername, printURL, printNote string

		if len(row.Username) >= 24 {
			printUsername = row.Username[:20]
			printUsername += "... "
		} else {
			printUsername = row.Username
		}
		if len(row.URL) >= 34 {
			printURL = row.URL[:30]
			printURL += "... "
		} else {
			printURL = row.URL
		}
		if len(row.Note) >= 40 {
			printNote = row.Note[:36]
			printNote += "... "
		} else {
			printNote = row.Note
		}

		fmt.Printf("%-5v %-25v %-35v %-40s\n", row.SID, printUsername, printURL, printNote)
	}
}

func pressEnterToContinue() {
	getStrInput("\n\nPress enter to continue...")
}

func printLogo() {

	fmt.Println(`
    ______             __     __                     __    __     
   /      \           /  |   /  |                   /  |  /  |    
  /$$$$$$  |  ______  $$ |   $$ | ______   __    __ $$ | _$$ |_   
  $$ | _$$/  /      \ $$ |   $$ |/      \ /  |  /  |$$ |/ $$   |  
  $$ |/    |/$$$$$$  |$$  \ /$$/ $$$$$$  |$$ |  $$ |$$ |$$$$$$/   
  $$ |$$$$ |$$ |  $$ | $$  /$$/  /    $$ |$$ |  $$ |$$ |  $$ | __ 
  $$ \__$$ |$$ \__$$ |  $$ $$/  /$$$$$$$ |$$ \__$$ |$$ |  $$ |/  |
  $$    $$/ $$    $$/    $$$/   $$    $$ |$$    $$/ $$ |  $$  $$/ 
   $$$$$$/   $$$$$$/      $/     $$$$$$$/  $$$$$$/  $$/    $$$$/  
					   `)
	fmt.Printf("                        %s\n\n", "https://GoVault.ch")

}

func printLogin() {

	fmt.Println(`
                 _                _        
                | |              (_)       
                | |  ___    __ _  _  _ __  
                | | / _ \  / _\ || || \_ \ 
                | || (_) || (_| || || | | |
                |_| \___/  \__/ ||_||_| |_|
                            __/ |          
                           |___/           												
					   `)

}

func printWelcome() {
	fmt.Println("Welcome to your GoVault " + dbGetDataBySID(0).Username + " !\n")
}

func printEventMessage(message string) {
	clearTerminal(true)
	fmt.Println(message)
	pressEnterToContinue()
}
