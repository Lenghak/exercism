package speed

// Car struct for car object
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		distance:     0,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

// Track struct for track object
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

	if car.battery < car.batteryDrain {
		return car
	}

	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	var driveCount int = car.battery / car.batteryDrain
	var totalDistance = (car.speed * driveCount) + car.distance

	return totalDistance >= track.distance
}
