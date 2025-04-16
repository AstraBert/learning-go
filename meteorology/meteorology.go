package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}


type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (t SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[t]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (t Speed) String() string {
	return fmt.Sprintf("%d %s", t.magnitude, t.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (t MeteorologyData) String() string {
	s := fmt.Sprintf("%s: %s, Wind %s at %s, %d", t.location, t.temperature.String(), t.windDirection, t.windSpeed.String(), t.humidity)
	return s+"% Humidity"
}