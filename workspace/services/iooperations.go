package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gablesiak/datatypes"
	"github.com/google/uuid"
)

func OpenFile() []byte {
	inputFile := "input.json"

	inputByte, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return inputByte
}


func CreateInputStruct() []datatypes.InputUserData{
	inputByte := OpenFile()

	inputStruct := []datatypes.InputUserData{}

	err := json.Unmarshal([]byte(inputByte), &inputStruct)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return inputStruct

}


func SaveUser(newUser datatypes.InputUserData){
	ValidateData(newUser)
	userEntry :=TransformData(newUser)
	GenerateOutput(userEntry)
}

func GenerateOutput(outputData datatypes.OutputUserData){
	multipleOutputData, err :=json.MarshalIndent(outputData, "", " ")
	if err != nil {
		fmt.Print(err)
	}
	
	uuidString := uuid.NewString()
	outputFile, err := os.Create("./output/" + uuidString + ".json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	_,_ = outputFile.Write(multipleOutputData)

}

