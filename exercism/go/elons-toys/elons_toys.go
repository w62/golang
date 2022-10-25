package elon
import (
	"fmt"
)

// TODO: define the 'Drive()' method

func (c * Car) Drive() {

	if c.battery - c.batteryDrain <= 0 {
		return
	}

	c.battery -= c.batteryDrain
	c.distance = c.distance + c.speed
}
// TODO: define the 'DisplayDistance() string' method

func (c Car) DisplayDistance () string {
	result := fmt.Sprintf("Driven %d meters", c.distance)
	return result
}

// TODO: define the 'DisplayBattery() string' method

func (c Car) DisplayBattery() string {
	result := fmt.Sprintf("Battery at %d%%", c.battery)
	return result
}

// TODO: define the 'CanFinish(trackDistance int) bool' method

func (c Car) CanFinish (trackDistance int) bool {
	return (c.battery / c.batteryDrain) *  c.speed >= trackDistance
}
