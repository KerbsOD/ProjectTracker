package app

import "time"

type Responsible interface {
	DaysToFinish(anEffort int) time.Duration
	AddWorkingDatesForEachDeveloper(aSliceOfContiguousDates []time.Time, aWorkingDatesArrayForEachDeveloper map[*Developer][]time.Time)
	AddResponsiblesTo(aCollector *[]Responsible)
}
