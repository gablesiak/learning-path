package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var multipleUserOutput []outputUserData

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		usersJson, err := json.Marshal(multipleUserOutput)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.Write(usersJson)

	case http.MethodPost:
		//var newOutput outputUserData
		var newInput inputUserData
		outputByte, err := ioutil.ReadAll(r.Body)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(outputByte, &newInput)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ValidateData(newInput)
		userEntry := TransformData(newInput)
		multipleUserOutput = append(multipleUserOutput,userEntry)
		w.WriteHeader(http.StatusCreated)
		return
		
}
}

func main() {

	multipleUserInput :=CreateInputStruct()


	for i :=0; i<len(multipleUserInput);i++{
		ValidateData(multipleUserInput[i])
		userEntry := TransformData(multipleUserInput[i])
		multipleUserOutput = append(multipleUserOutput,userEntry)
	}

	GenerateOutput(multipleUserOutput)

	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":5000", nil)
}
