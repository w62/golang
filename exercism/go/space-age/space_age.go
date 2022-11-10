package space
import (
	"fmt"
)

type Planet string

const earthyear = 31557600 

func Age(seconds float64, planet Planet) float64 {
	fmt.Println(planet)
	var factor float64 
	switch planet {
	case "Earth":
		factor = 1.0
		case "Mercury":
			factor = 0.2408467
		case "Venus":
			factor = 0.61519726
		case "Mars":
			factor = 1.8808158
		case "Jupiter":
			factor = 11.862615
		case "Saturn":
			factor = 29.447498
		case "Uranus":
			factor = 84.016846
		case "Neptune":
			factor = 164.79132
		default:
			return -1.0
}
	return seconds / (factor * earthyear)
	panic("Please implement the Age function")
}
