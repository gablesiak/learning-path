package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gablesiak/datatypes"
	"github.com/google/uuid"
)

func SaveUser(newUser datatypes.InputUserData) {
	validateInputStruct(newUser)
	userEntry := transformData(newUser)
	generateOutput(userEntry)
}

func GenerateOutputStruct(newUser datatypes.InputUserData) datatypes.OutputUserData {
	validateInputStruct(newUser)
	userEntry := transformData(newUser)

	return userEntry
}

func generateOutput(outputData datatypes.OutputUserData) {
	multipleOutputData, err := json.MarshalIndent(outputData, "", " ")
	if err != nil {
		fmt.Print(err)
	}

	uuidString := uuid.NewString()
	outputFile, err := os.Create("./output/" + uuidString + ".json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	_, err = outputFile.Write(multipleOutputData)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
