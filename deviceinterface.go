package main

type DeviceConfig struct {
	Executable         string  `json:"executable"`
	DeviceType         string  `json:"deviceType"`
	ColorHue           float64 `json:"ColorHue"`
	ColorHueMultiplier float64 `json:"ColorHueMultiplier"`
	ColorSaturation    float64 `json:"ColorSaturation"`
	ColorValue         float64 `json:"ColorValue"`
	DefaultSymbol      string  `json:"defaultSymbol"`
	DefaultDeviceName  string  `json:"defaultDeviceName"`
	DeviceHostIP       string  `json:"DeviceHostIP"`
}

// View ... This interface allows for reacting to events
type Device interface {
	Render(string, *DeviceConfig) bool
}
