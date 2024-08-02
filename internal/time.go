package internal

import "time"

const Day = 24 * time.Hour

func GreaterDuration(a, b time.Duration) bool {
	return a > b
}

func MaxDateBetween(aDate, anotherDate time.Time) time.Time {
	return Max(aDate, anotherDate, time.Time.After) // Method expression
}
