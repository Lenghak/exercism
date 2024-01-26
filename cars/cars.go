package cars

// IndividualProduceCost - the cost to produce an individual car
const IndividualProduceCost = 10000

// GroupOfTenProduceCost - the cost to produce a group of ten cars
const GroupOfTenProduceCost = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupOfTen int = carsCount / 10
	var remainder int = carsCount % 10

	return uint(groupOfTen)*GroupOfTenProduceCost + uint(remainder)*IndividualProduceCost
}
