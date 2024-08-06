package pkg

import "time"

type Task interface {
	StartDate() time.Time
	FinishDate() time.Time
	AddWorkingDatesByDeveloperTo(map[*Developer][]time.Time)
}
