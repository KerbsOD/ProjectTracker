package pkg

import (
	"Project/internal"
	"math"
	"time"
)

type Developer struct {
	name       string
	dedication int
	rate       int
}

func NewDeveloper(name string, dedication int, rate int) *Developer {
	d := new(Developer)
	d.name = name
	d.dedication = dedication
	d.rate = rate
	return d
}

func (d Developer) DaysToFinish(effort int) time.Duration {
	workSessions := int(math.Ceil(float64(effort) / float64(d.dedication)))
	fullWorkDays := time.Duration(workSessions) * internal.Day
	return fullWorkDays
}

func (d *Developer) AddDatesToDeveloper(datesPerDeveloper map[*Developer][]time.Time, workingDates []time.Time) {
	if _, ok := datesPerDeveloper[d]; !ok {
		datesPerDeveloper[d] = []time.Time{}
	}
	datesPerDeveloper[d] = append(datesPerDeveloper[d], workingDates...)
}
