package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/go-playground/validator.v9"
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

func OpenOutputFile() []byte {
	outputFile := "output.json"

	outputByte, err := ioutil.ReadFile(outputFile)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return outputByte
}

func CreateInputStruct() []inputUserData{
	inputByte := OpenFile()

	inputStruct := []inputUserData{}

	err := json.Unmarshal([]byte(inputByte), &inputStruct)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return inputStruct

}

func ValidateData(inputStruct inputUserData) {
	validateInput := validator.New()

	err := validateInput.Struct(inputStruct)

	if err != nil || strings.Count(inputStruct.Organization, "/") != 3 {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func TransformData(inputStruct inputUserData) outputUserData {
	splitedInput, department, subdepartment, team := splitOrganizationString(inputStruct)

	return outputUserData{
		FullName:      inputStruct.FirstName + " " + inputStruct.LastName,
		Age:           inputStruct.Age,
		City:          inputStruct.City,
		Organization:  splitedInput,
		Department:    department,
		Subdepartment: subdepartment,
		Team:          team,
	}
}

func splitOrganizationString(inputStruct inputUserData) (string, string, string, string) {
	splitedInput := strings.Split(inputStruct.Organization, "/")
	return splitedInput[0], splitedInput[1], splitedInput[2], splitedInput[3]
}


func GenerateOutput(outputData []outputUserData){
	multipleOutputData, err :=json.MarshalIndent(outputData, "", " ")
	if err != nil {
		fmt.Print(err)
	}
	
	fmt.Println(string(multipleOutputData))

	outPutFile, err := os.Create("./output.json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	outPutFile.Write(multipleOutputData)

}

