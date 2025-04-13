package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for i := 0; i < len(birdsPerDay) ; i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekBirds := []int{}	
	if (week)*7 <= len(birdsPerDay) {
		weekBirds = birdsPerDay[(week-1)*7:(week)*7]
		var weekSum int
		for i := 0; i < len(weekBirds); i++{
			weekSum += weekBirds[i]
		}
		return weekSum
	} else {
		return 0
	}
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i % 2 == 0 {
			birdsPerDay[i] = birdsPerDay[i]+1
		} else {
			birdsPerDay[i] = birdsPerDay[i]
		}
	}
	return birdsPerDay
}
