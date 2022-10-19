package purchase

import (
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	default:
		return false
	}
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var choice string
	if strings.Compare(option1, option2) < 1 {
		choice = option1
	} else {
		choice = option2
	}

	return choice + " is clearly the better choice."
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	myAge := int(age)

	if myAge < 3 {
		return originalPrice * 0.8 
	} else {
		if myAge >= 3 && myAge < 10 {
			return originalPrice * 0.7
		} else {
			return originalPrice * 0.5
		}
	}
	panic("CalculateResellPrice not implemented")
}
