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

const jsonFilePath = "GoVault.json"

func dbRead() []PWData {
	//Open File
	file, _ := os.Open(jsonFilePath)
	defer file.Close()
	//Read Data
	data, _ := io.ReadAll(file)
	var pwDataList []PWData
	//Get Infos
	json.Unmarshal(data, &pwDataList)
	return pwDataList
}

func dbAddData(newData []PWData) {
	//Get Data
	existingList := dbRead()
	// Append the new data to the existing list
	existingList = append(existingList, newData...)
	dbSave(existingList)
}

func dbDelete(delSID int) {
	//Del admin protection
	if delSID == 0 {
		return
	}
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	var newPwDataList []PWData
	for _, pwData := range pwDataList {

		if delSID != pwData.SID {
			newPwDataList = append(newPwDataList, pwData)
		}
	}
	//Write new DB
	if len(pwDataList) == len(newPwDataList) {
		fmt.Println("No data to delete!")
		return
	}
	dbSave(newPwDataList)

	//fmt.Println("The data has been successfully deleted.")
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

func dbGetDataBySearch(search string) []PWData {
	//get DB
	pwDataList := dbRead()
	//get all exept delID
	search = strings.ToLower(search)
	var returnPWData []PWData
	for _, pwData := range pwDataList {

		if pwData.SID != 0 && strings.Contains(strings.ToLower(pwData.Username), search) || strings.Contains(strings.ToLower(pwData.URL), search) || strings.Contains(strings.ToLower(pwData.Note), search) {
			returnPWData = append(returnPWData, pwData)
		}
	}
	return returnPWData
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
	dbSave(newpwDataList)
}

// save database
func dbSave(newpwDataList []PWData) {
	updatedData, _ := json.MarshalIndent(newpwDataList, "", "  ")
	os.WriteFile(jsonFilePath, updatedData, 0644)
}
