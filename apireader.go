package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var apiconfigpath = "apiconfig.json"

type ApiConfig struct {
	Regex        string
	BaseURL      string
	Apikey       string
	DataEndpoint string
	UserAgent    string
}

// Read the API client configuration json file and populate ApiConfig struct
func loadAPIConfiguration() ApiConfig {

	// Read the Config file
	jsonFile, err := ioutil.ReadFile(apiconfigpath)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: Could not read from file: %s, err: %s", apiconfigpath, err))
	}

	//Declare a map variable
	apiconfig := ApiConfig{}

	// Read file as byte array
	err = json.Unmarshal([]byte(jsonFile), &apiconfig)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: Could not Unmarshall json: %s, err: %s", apiconfigpath, err))
	}

	fmt.Println(fmt.Sprintf("INFO: Loaded API configuration: %s\n", apiconfig))

	return apiconfig
}
