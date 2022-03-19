package main

import (
	"fmt"
	"strconv"
)

type SimpleDevice struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Bold   string
}

func NewSimpleDevice() *SimpleDevice {
	sd := new(SimpleDevice)
	sd.Reset = "\033[0m"
	sd.Red = "\033[31m"
	sd.Green = "\033[32m"
	sd.Yellow = "\033[33m"
	sd.Bold = "\033[1m"

	return sd
}

// Define terminal colors

// Implement the Render function
func (d *SimpleDevice) Render(input string, dc *DeviceConfig) bool {

	result, error := strconv.ParseFloat(input, 64)
	if error != nil {
		fmt.Println(fmt.Sprintf("ERROR: Error converting change percentage: %2.f, Error: %s", result, error))
		return false
	}

	color := d.Green

	if result < 0 {
		color = d.Red
	}

	fmt.Println(fmt.Sprintf("%s %s %.2f %% %s", d.Bold, color, result, d.Reset))

	return true

}
