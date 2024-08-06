package pkg

import "time"

type Responsible interface {
	DaysToFinish(effort int) time.Duration
	AddDatesToDeveloper(datesPerDeveloper map[*Developer][]time.Time, workingDates []time.Time)
}
