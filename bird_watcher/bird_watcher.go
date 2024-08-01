package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	result := 0

	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	startDay := 0

	if 7*(week-1) > 0 {
		startDay = 7 * (week - 1)
	}

	if 7*(week-1) > len(birdsPerDay) {
		return 0
	}

	endDay := len(birdsPerDay)

	if startDay+7 < endDay {
		endDay = startDay + 7
	}

	birdsWeek := birdsPerDay[startDay:endDay]
	return TotalBirdCount(birdsWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
