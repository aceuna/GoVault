package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func clearTerminal() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system")
	}
	printLogo()
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func generateNewSID() int {
	//get DB
	pwDataList := dbRead()
	//SID +1
	return pwDataList[len(pwDataList)-1].SID + 1
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

}

func doesSidExist(searchSID PWData) bool {

	var noData PWData

	if noData != searchSID {

		return true

	}

	return false
}

func checkPwWithHash() (string, bool) {
	var pw string

	pw = getStrInput("Enter your PW")

	if createMD5Hash(pw) == dbGetDataBySID(0).Password {
		return pw, true
	}
	return pw, false
}
