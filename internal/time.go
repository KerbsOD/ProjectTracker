package internal

import "time"

const Day = 24 * time.Hour

func GreaterDuration(a, b time.Duration) bool {
	return a > b
}

func MaxDateBetween(aDate, anotherDate time.Time) time.Time {
	return Max(aDate, anotherDate, time.Time.After) // Method expression
}

func MaxDateInArray(array []time.Time) time.Time {
	return MaximizeElementByComparer(array, time.Time.After)
}

func MinDateInArray(array []time.Time) time.Time {
	return MaximizeElementByComparer(array, time.Time.Before)
}
