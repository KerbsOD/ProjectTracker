package pkg

import (
	"Project/internal"
	"errors"
	"math"
	"time"
)

type Developer struct {
	name       string
	dedication int
	rate       int
}

func NewDeveloper(aName string, aDedication int, aRate int) *Developer {
	assertValidDeveloper(aName, aDedication, aRate)
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

func (d *Developer) AddResponsiblesTo(aCollector *[]Responsible) {
	*aCollector = append(*aCollector, d)
}

func (d Developer) CostForWorking(aNumberOfDays int) int {
	return d.rate * d.dedication * aNumberOfDays
}

/*
	PRIVATE
*/

func assertValidDeveloper(aName string, aDedication int, aRate int) {
	assertValidDeveloperName(aName)
	assertValidDedication(aDedication)
	assertValidRate(aRate)
}

func assertValidDeveloperName(aName string) {
	if internal.EmptyName(aName) {
		panic(errors.New(internal.InvalidDeveloperNameErrorMessage))
	}
}

func assertValidDedication(aDedication int) {
	if aDedication < 1 {
		panic(errors.New(internal.InvalidDeveloperDedicationErrorMessage))
	}
}

func assertValidRate(aRate int) {
	if aRate < 1 {
		panic(errors.New(internal.InvalidDeveloperRateErrorMessage))
	}
}
