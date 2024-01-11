package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func clearTerminal(logo bool) {
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
	if logo {
		printLogo()
	}
}

/*
	func isNumeric(s string) bool {
		_, err := strconv.ParseFloat(s, 64)
		return err == nil
	}
*/
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

func doesSidExist(searchSID PWData) bool {
	var noData PWData
	return noData != searchSID

}

func checkPwWithHash() (string, bool) {

	pw := getStrInput("Enter your GoVault password")

	if createMD5Hash(pw) == dbGetDataBySID(0).Password {
		return pw, true
	}
	return pw, false
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func exit() {
	clearTerminal(false)
	os.Exit(1)
}
