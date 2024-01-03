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
	strNum := strconv.Itoa(data.SID)

	fmt.Println("SID:", strNum)
	fmt.Println("Username:", data.Username)
	fmt.Println("Password:", pw)
	fmt.Println("URL:", data.URL)
	fmt.Println("Note:", data.Note)
	fmt.Println()
}

func newDecision(question string) string {
	var decision string
	for {
		decision = getStrInput(question)
		decision = strings.ToLower(decision)
		if decision == "y" || decision == "n" {
			break
		}
	}
	return decision
}

func formatTable(PWData []PWData) {
	// Print table header
	fmt.Printf("%-5s %-30s %-40s %-15s\n", "SID", "USERNAME", "URL", "NOTE")
	// Print table rows
	for _, row := range PWData {
		fmt.Printf("%-5v %-30v %-40v %-15s\n", row.SID, row.Username, row.URL, row.Note)
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
