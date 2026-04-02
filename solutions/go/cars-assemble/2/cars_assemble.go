package cars

const percentBase = 100
const minutesInHours = 60
const groupSize = 10
const oneCarCost = 10000
const groupOfCarsCost = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate / percentBase * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(successRate/percentBase*float64(productionRate)) / minutesInHours
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint(carsCount/groupSize*groupOfCarsCost + carsCount%groupSize*oneCarCost)
}
