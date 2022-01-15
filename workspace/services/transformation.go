package services

import (
	"strings"

	"github.com/gablesiak/datatypes"
)

func transformData(inputStruct datatypes.InputUserData) datatypes.OutputUserData {
	splitedInput, department, subdepartment, team := splitOrganizationString(inputStruct)

	return datatypes.OutputUserData{
		FullName:      inputStruct.FirstName + " " + inputStruct.LastName,
		Age:           inputStruct.Age,
		City:          inputStruct.City,
		Organization:  splitedInput,
		Department:    department,
		Subdepartment: subdepartment,
		Team:          team,
		
	}
}

func splitOrganizationString(inputStruct datatypes.InputUserData) (string, string, string, string) {
	splitedInput := strings.Split(inputStruct.Organization, "/")
	return splitedInput[0], splitedInput[1], splitedInput[2], splitedInput[3]
}