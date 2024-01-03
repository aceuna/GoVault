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
func getIntInput(inputMessage string) int {

	for {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print(inputMessage + ": ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}

		userInput = strings.TrimSpace(userInput)

		if isInteger(userInput) {
			intInput, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				log.Fatal()
			}

			return intInput
		}

	}
}

func formatPWData(data PWData) {

	strNum := strconv.Itoa(data.SID)

	fmt.Println(strNum, data.Username, data.URL, data.Note)

}

func formatPWDataWithPW(data PWData, pw string) {

	strNum := strconv.Itoa(data.SID)

	fmt.Println("SID:", strNum)
	fmt.Println("Username:", data.Username)
	fmt.Println("Password:", pw)
	fmt.Println("URL:", data.URL)
	fmt.Println("Note:", data.Note)
	fmt.Println("\n")

}

func getDecision(question string) string {
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

func printTable(PWData []PWData) {
	// Print table header
	fmt.Printf("%-5s %-30s %-40s %-15s\n", "SID", "Username", "URL", "Note")
	// Print table rows
	for _, row := range PWData {
		fmt.Printf("%-5v %-30v %-40v %-15s\n", row.SID, row.Username, row.URL, row.Note)
	}
}

func pressEnter() {
	getStrInput("\n\nPress enter to continue...")
}
