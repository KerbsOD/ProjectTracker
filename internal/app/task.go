package app

import "time"

type Task interface {
	StartDate() time.Time
	FinishDate() time.Time
	AddWorkingDatesForEachDeveloper(aWorkingDatesSliceForEachDeveloper map[*Developer][]time.Time)
}
