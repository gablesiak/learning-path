package main

import (
	"encoding/json"
	"fmt"
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
	inputByte, err := ioutil.ReadFile("input.json")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return inputByte
}

func CreateInputStruct() inputUserData {
	inputByte := OpenFile()

	inputStruct := inputUserData{}

	_ = json.Unmarshal([]byte(inputByte), &inputStruct)

	return inputStruct
}

func ValidateData() {

	inputStruct := CreateInputStruct()

	validateInput := validator.New()

	err := validateInput.Struct(inputStruct)

	if err != nil || strings.Count(inputStruct.Organization, "/") != 3 {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Validation done, no issues found")
	}
}

func TransformData() outputUserData {

	inputStruct := CreateInputStruct()

	ValidateData()

	outputStruct := outputUserData{}

	splitedInput := strings.Split(inputStruct.Organization, "/")

	outputStruct.FullName = inputStruct.FirstName + " " + inputStruct.LastName
	outputStruct.Age = inputStruct.Age
	outputStruct.City = inputStruct.City
	outputStruct.Organization = splitedInput[0]
	outputStruct.Department = splitedInput[1]
	outputStruct.Subdepartment = splitedInput[2]
	outputStruct.Team = splitedInput[3]

	return outputStruct
}

func GenerateOutput(outputData outputUserData) {
	outputFile, _ := json.MarshalIndent(outputData, "", " ")

	_ = ioutil.WriteFile("output.json", outputFile, 0644)
}

func main() {

	userData := TransformData()

	GenerateOutput(userData)
}
