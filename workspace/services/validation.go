package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gablesiak/datatypes"
	"gopkg.in/go-playground/validator.v9"
)

func validateData(inputStruct datatypes.InputUserData) {
	validateInput := validator.New()

	err := validateInput.Struct(inputStruct)

	if err != nil {
		fmt.Println(err.Error())
	}

	validateOrganization, err := regexp.MatchString("[^A-Za-z/A-Za-z/A-Za-z/A-Za-z]", inputStruct.Organization)

	if err != nil || !validateOrganization{
		fmt.Println(err.Error())
	}
}


func ValidateInputStructure(r *http.Request) (datatypes.InputUserData, error){
	var newUser datatypes.InputUserData
	outputByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return newUser, err
	}
	
	err = json.Unmarshal(outputByte, &newUser)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
