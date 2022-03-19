package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type ColorBulbDevice struct{}

func NewColorBulbDevice() *ColorBulbDevice {
	cd := new(ColorBulbDevice)

	return cd
}

func ExecStartAndWait(cmd *exec.Cmd) bool {
	success := true

	if Debug {
		fmt.Println(fmt.Sprintf("DEBUG: Trying command: ", cmd))
	}

	stderr, err1 := cmd.StderrPipe()
	if err1 != nil {
		log.Fatal(err1)
		success = false
	}

	stdout, err2 := cmd.StdoutPipe()
	if err2 != nil {
		log.Fatal(err2)
		success = false
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(fmt.Sprintf("WARN: An error occurred when launching process %s: %s", cmd, err))
		success = false
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)
	success = !strings.Contains(strings.ToLower(string(slurp)), "error:")

	slurp, _ = io.ReadAll(stdout)
	fmt.Printf("%s\n", slurp)

	//if err := cmd.Wait(); err != nil {
	//	log.Fatal(err)
	//	success = false
	//}

	return success

}

// Implement the Render function
func (d *ColorBulbDevice) Render(input string, dc *DeviceConfig) bool {

	success := true

	result, error := strconv.ParseFloat(input, 64)
	if error != nil {
		fmt.Println(fmt.Sprintf("ERROR: Error converting change percentage: %2.f, Error: %s", result, error))
		return false
	}

	result = dc.ColorHue + result*dc.ColorHueMultiplier
	h := fmt.Sprintf("%.0f", result)
	s := fmt.Sprintf("%.0f", dc.ColorSaturation)
	v := fmt.Sprintf("%.0f", dc.ColorValue)

	if Debug {
		fmt.Println(fmt.Sprintf("DEBUG: Constructed Color string: %s %s %s", h, s, v))
	}

	var cmd, cmd2 *exec.Cmd
	// First exec try
	//commandargs := fmt.Sprintf("%s %s %s %s %s %s", "--type", dc.DeviceType, "--alias", dc.DefaultDeviceName, "hsv", colorstring)
	if len(dc.DeviceHostIP) > 0 {
		cmd = exec.Command(dc.Executable, "--type", dc.DeviceType, "--host", dc.DeviceHostIP, "hsv", h, s, v)
	} else {
		cmd = exec.Command(dc.Executable, "--type", dc.DeviceType, "--alias", dc.DefaultDeviceName, "hsv", h, s, v)
	}

	if !ExecStartAndWait(cmd) {
		fmt.Println(fmt.Sprintf("INFO: Retrying with older command format"))
		success = false

		// Second exec try
		devicestr := fmt.Sprintf("--%s", dc.DeviceType)
		//commandargs = fmt.Sprintf("%s %s %s %s %s", devicestr, "--alias", dc.DefaultDeviceName, "hsv", colorstring)
		if len(dc.DeviceHostIP) > 0 {
			cmd2 = exec.Command(dc.Executable, devicestr, "--host", dc.DeviceHostIP, "hsv", h, s, v)
		} else {
			cmd2 = exec.Command(dc.Executable, devicestr, "--alias", dc.DefaultDeviceName, "hsv", h, s, v)
		}

		success = ExecStartAndWait(cmd2)
	}

	return success

}
