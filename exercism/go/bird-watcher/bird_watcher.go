package birdwatcher

import (
	"fmt"
)

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var result int
	for i := 0; i<len(birdsPerDay); i++ {
		result = result + birdsPerDay[i]
	}
	return result
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var result int

	start := (week - 1) * 7

	fmt.Println("week=", week)
	
	for i := start ; i< (start) + 7 ; i++ {
		fmt.Println("i=", i, "birdsPerDay=", birdsPerDay[i])
		result = result + birdsPerDay[i]
	}
		fmt.Println("result=", result)
	return result
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i :=0; i<len(birdsPerDay); i ++ {
		if i%2==0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
