package pkg

import "time"

type Responsible interface {
	DaysToFinish(effort int) time.Duration
}
