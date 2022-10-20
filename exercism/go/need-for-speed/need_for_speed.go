package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}

type Track struct {
	distance int
}

var (
	battery int
	batteryDrain int
	speed int
	distance int
)

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	var value Car
	value.speed = speed
	value.batteryDrain = batteryDrain
	value.battery = 100
	value.distance = 0

	return value
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	var result Track
	result.distance = distance
	return result
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
	car.battery = car.battery - car.batteryDrain
	car.distance = car.distance + car.speed
}
	return car
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return (car.battery / car.batteryDrain) *  car.speed >= track.distance
	panic("Please implement the CanFinish function")
}
