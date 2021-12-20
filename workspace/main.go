package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type inputUserData struct {
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	Age          int    `validate:"required,gte=18,lte=80"`
	City         string `validate:"required"`
	Organization string `validate:"required,contains=/"`
}

type outputUserData struct {
	FullName      string
	Age           int
	City          string
	Organization  string
	Department    string
	Subdepartment string
	Team          string
}

func OpenFile() []byte {
	inputFile := "input.json"

	inputByte, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return inputByte
}

func CreateInputStruct() inputUserData {
	inputByte := OpenFile()

	inputStruct := inputUserData{}

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

func GenerateOutput(outputData outputUserData) {
	outputFile, err := json.MarshalIndent(outputData, "", " ")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	//bit flags for output file - Owner: read & write Group: read Other: read
	var fileMode fs.FileMode = 0644
	_ = ioutil.WriteFile("output.json", outputFile, fileMode)
}

func main() {

	inputUserData := CreateInputStruct()
	ValidateData(inputUserData)
	outputUserData := TransformData(inputUserData)
	GenerateOutput(outputUserData)
}
