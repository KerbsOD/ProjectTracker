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

func NewDeveloper(aName string, aDedication int, aRate int) *Developer {
	d := new(Developer)
	d.name = aName
	d.dedication = aDedication
	d.rate = aRate
	return d
}

func (d Developer) DaysToFinish(anEffort int) time.Duration {
	workSessions := int(math.Ceil(float64(anEffort) / float64(d.dedication)))
	fullWorkDays := time.Duration(workSessions) * internal.Day
	return fullWorkDays
}

func (d *Developer) AddWorkingDatesForEachDeveloper(aSliceOfContiguousDates []time.Time, aWorkingDatesArrayForEachDeveloper map[*Developer][]time.Time) {
	if _, ok := aWorkingDatesArrayForEachDeveloper[d]; !ok {
		aWorkingDatesArrayForEachDeveloper[d] = []time.Time{}
	}
	aWorkingDatesArrayForEachDeveloper[d] = append(aWorkingDatesArrayForEachDeveloper[d], aSliceOfContiguousDates...)
}
