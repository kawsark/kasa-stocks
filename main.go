package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Debug ... turn on or off Debug mode
var Debug, NetDebug bool

// Configuration variables
var configFile, configFileSet = os.LookupEnv("CONFIG_FILE")
var refreshInterval int
var devicename, devicetype, symbol, defaultHSV string
var EffectiveDeviceConfig DeviceConfig

// Read the configuration json file and populate DeviceConfig struct
func loadDeviceConfiguration() DeviceConfig {
	// Set a default configfile if its not set already
	if !configFileSet {
		configFile = "config.json"
	}

	// Read the Config file
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: Could not read from file: %s, err: %s", configFile, err))
	}

	//Declare a map variable
	EffectiveDeviceConfig = DeviceConfig{}

	// Read file as byte array
	err = json.Unmarshal([]byte(jsonFile), &EffectiveDeviceConfig)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: Could not Unmarshall json: %s, err: %s", configFile, err))
	}

	return EffectiveDeviceConfig
}

func init() {
	// Lookup DEBUG environment variable and set it globally
	v := os.Getenv("DEBUG")
	var err error
	Debug, err = strconv.ParseBool(v)
	if err != nil {
		Debug = false
	}

	// Lookup DEBUG environment variable and set it globally
	n := os.Getenv("NET_DEBUG")
	NetDebug, err = strconv.ParseBool(n)
	if err != nil {
		NetDebug = false
	}

	config := loadDeviceConfiguration()

	// Set flags (overriding the file configuration if it exists)
	flag.IntVar(&refreshInterval, "n", -1, "Specifies the number of seconds to lookup stock price")
	flag.StringVar(&devicename, "name", config.DefaultDeviceName, "Specifies the target device")
	flag.StringVar(&devicetype, "type", config.DeviceType, "Specifies the target device type")
	flag.StringVar(&symbol, "symbol", config.DefaultSymbol, "Specifies the symbol to lookup")

}

func main() {
	flag.Parse()

	fmt.Println("INFO: Loaded configuration: ", EffectiveDeviceConfig)

	var device Device

	if devicetype == "colorbulb" {
		device = NewColorBulbDevice()
	} else {
		device = NewSimpleDevice()
	}

	changePercent := readChangePercent(symbol)
	device.Render(changePercent, &EffectiveDeviceConfig)
}
