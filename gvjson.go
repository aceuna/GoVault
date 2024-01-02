package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type PWData struct {
	SID      int    `json:"sid"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Note     string `json:"note"`
}

var filePath = "GoVault.json"

func dbRead() []PWData {
	//Open File
	file, _ := os.Open(filePath)
	defer file.Close()
	//Read Data
	data, _ := io.ReadAll(file)
	var pwDataList []PWData
	//Get Infos
	json.Unmarshal(data, &pwDataList)
	return pwDataList

}

func dbAppend(newData []PWData) {
	//Get Data
	existingList := dbRead()
	// Append the new data to the existing list
	existingList = append(existingList, newData...)
	updatedData, _ := json.MarshalIndent(existingList, "", "  ")
	os.WriteFile(filePath, updatedData, 0644)
	fmt.Println("Data appended and updated successfully.")
}

func dbDelete(delID int) {
	//Del admin protection
	if delID == 0 {
		fmt.Println("Do not delete the admin!")
		return
	}
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	var newpwDataList []PWData
	for _, pwData := range pwDataList {

		if delID != pwData.SID {
			newpwDataList = append(newpwDataList, pwData)
		}
	}
	//Write new DB

	if len(pwDataList) == len(newpwDataList) {
		fmt.Println("No data to delete!")
		return
	}
	updatedData, _ := json.MarshalIndent(newpwDataList, "", "  ")
	os.WriteFile(filePath, updatedData, 0644)

	fmt.Println("The data has been successfully deleted.")
}

func dbGetALLData() {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	fmt.Println("SID, User, URL, Note")
	for _, pwData := range pwDataList {
		if pwData.SID != 0 {
			formatPWData(pwData)
		}
	}

}

func dbGetDataBySID(searchSID int) PWData {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	for _, pwData := range pwDataList {

		if searchSID == pwData.SID {
			return pwData
		}
	}

	return PWData{}
}

func dbGetSearch(search string) []PWData {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	search = strings.ToLower(search)
	var returnPWData []PWData
	for _, pwData := range pwDataList {

		if pwData.SID != 0 && strings.Contains(strings.ToLower(pwData.Username), search) || strings.Contains(strings.ToLower(pwData.URL), search) || strings.Contains(strings.ToLower(pwData.Note), search) {

			//formatPWData(dbGetDataBySID(pwData.SID))
			returnPWData = append(returnPWData, pwData)

		}
	}
	return returnPWData
}

func dbGetSearchReturn(search string) PWData {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	search = strings.ToLower(search)
	for _, pwData := range pwDataList {

		if pwData.SID != 0 && strings.Contains(strings.ToLower(pwData.Username), search) || strings.Contains(strings.ToLower(pwData.URL), search) || strings.Contains(strings.ToLower(pwData.Note), search) {

			fmt.Println(dbGetDataBySID(pwData.SID))

		}
	}
	return PWData{}
}

func dbAddPassword(searchSID int) PWData {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	for _, pwData := range pwDataList {

		if searchSID == pwData.SID {
			return pwData
		}
	}

	return PWData{}
}

func dbReplaceData(data PWData) {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	var newpwDataList []PWData
	for _, pwData := range pwDataList {

		if pwData.SID == data.SID {

			newpwDataList = append(newpwDataList, data)

		} else {
			newpwDataList = append(newpwDataList, pwData)
		}
	}

	updatedData, _ := json.MarshalIndent(newpwDataList, "", "  ")
	os.WriteFile(filePath, updatedData, 0644)
	fmt.Println("Data updated successfully.")
}
