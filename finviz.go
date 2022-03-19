package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var apiKey, baseURL, symbolEndpoint, changeRegex, userAgent string

//TODO: Populate default values if there was an issue
func init() {
	apiconfig := loadAPIConfiguration()
	apiKey = apiconfig.Apikey
	baseURL = apiconfig.BaseURL
	symbolEndpoint = apiconfig.DataEndpoint
	changeRegex = apiconfig.Regex
	userAgent = apiconfig.UserAgent
}

// Reads the specified URL by prepending the base URL and appending the token
// Returns result as a String
// TODO: Add retry logic
func readAPI(symbol string) string {

	// This is the request URL with the API token appended
	url := fmt.Sprint(baseURL, fmt.Sprintf(symbolEndpoint, symbol))

	fmt.Println(fmt.Sprintf("INFO: Request URL:  -> %s", url))

	// New client
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Chrome")

	response, err := client.Do(req)

	if err != nil {
		log.Println("ERROR: Received error from HTTP GET request", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ERROR: Received error from reading response", err)
	}

	resultStr := fmt.Sprintf(string(body))
	if NetDebug {
		fmt.Println(fmt.Sprintf("DEBUG: Response:  -> %s", resultStr))
	}

	return resultStr
}

// Returns the percentage
func readChangePercent(symbol string) string {
	apiResponse := readAPI(symbol)
	regularExpression, _ := regexp.Compile(changeRegex)

	matches := regularExpression.FindStringSubmatch(apiResponse)
	result := ""
	if len(matches) == 0 {
		fmt.Println(fmt.Sprintf("WARN: Cound not find a Change percentage with Regex: %s, please doublecheck URL or Regex.", changeRegex))
	} else {
		result = matches[len(matches)-1]
	}

	if Debug {
		fmt.Println(fmt.Sprintf("DEBUG: Returning Regex match:  -> %s", result))
	}

	return result
}
