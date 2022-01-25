package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gablesiak/datatypes"
	"github.com/google/uuid"
)

func GenerateLocalOutput(newUser datatypes.InputUserData) {
	userEntry := GenerateOutputStruct(newUser)
	generateLocalFile(userEntry)
}


func generateLocalFile(outputData datatypes.OutputUserData) {
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
