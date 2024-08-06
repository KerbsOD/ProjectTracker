package pkg

import (
	"Project/internal"
	"time"
)

type Team struct {
	name         string
	responsibles []Responsible
}

func NewTeam(aName string, aResponsibles []Responsible) *Team {
	t := new(Team)
	t.name = aName
	t.responsibles = aResponsibles
	return t
}

func (t Team) DaysToFinish(anEffort int) time.Duration {
	daysToCompleteTaskForEachResponsible := t.daysToCompleteTaskForEachResponsible(anEffort)
	maxDaysToCompleteTaskAmongResponsibles := internal.MaximizeElementByComparer(daysToCompleteTaskForEachResponsible, internal.GreaterDuration)
	return maxDaysToCompleteTaskAmongResponsibles
}

func (t Team) daysToCompleteTaskForEachResponsible(anEffort int) []time.Duration {
	daysToCompleteForEachResponsible := internal.Map(t.responsibles, func(aResponsible Responsible) time.Duration { return aResponsible.DaysToFinish(anEffort) })
	return daysToCompleteForEachResponsible
}

func (t Team) AddWorkingDatesForEachDeveloper(aSliceOfContiguousDates []time.Time, aWorkingDatesArrayForEachDeveloper map[*Developer][]time.Time) {
	for _, responsible := range t.responsibles {
		responsible.AddWorkingDatesForEachDeveloper(aSliceOfContiguousDates, aWorkingDatesArrayForEachDeveloper)
	}
}
