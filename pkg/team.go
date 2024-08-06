package pkg

import (
	"Project/internal"
	"time"
)

type Team struct {
	name         string
	responsibles []Responsible
}

func NewTeam(name string, members []Responsible) *Team {
	t := new(Team)
	t.name = name
	t.responsibles = members
	return t
}

func (t Team) DaysToFinish(effort int) time.Duration {
	daysToCompleteTaskForEachResponsible := t.daysToCompleteTaskForEachResponsible(effort)
	maxDaysToCompleteTaskAmongResponsibles := internal.MaximizeElementByComparer(daysToCompleteTaskForEachResponsible, internal.GreaterDuration)
	return maxDaysToCompleteTaskAmongResponsibles
}

func (t Team) daysToCompleteTaskForEachResponsible(effort int) []time.Duration {
	daysToCompleteForEachResponsible := internal.Map(t.responsibles, func(aResponsible Responsible) time.Duration { return aResponsible.DaysToFinish(effort) })
	return daysToCompleteForEachResponsible
}

func (t Team) AddDatesToDeveloper(datesPerDeveloper map[*Developer][]time.Time, workingDates []time.Time) {
	for _, responsible := range t.responsibles {
		responsible.AddDatesToDeveloper(datesPerDeveloper, workingDates)
	}
}
