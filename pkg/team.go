package pkg

import (
	"Project/internal"
	"errors"
	"strings"
	"time"
)

type Team struct {
	name         string
	responsibles []Responsible
}

func NewTeam(aName string, aResponsibles []Responsible) *Team {
	assertValidTeam(aName, aResponsibles)
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

func (t Team) AddWorkingDatesForEachDeveloper(aSliceOfContiguousDates []time.Time, aWorkingDatesArrayForEachDeveloper map[*Developer][]time.Time) {
	for _, responsible := range t.responsibles {
		responsible.AddWorkingDatesForEachDeveloper(aSliceOfContiguousDates, aWorkingDatesArrayForEachDeveloper)
	}
}

func (t Team) AddResponsiblesTo(aCollector *[]Responsible) {
	*aCollector = append(*aCollector, &t)
	for _, responsible := range t.responsibles {
		responsible.AddResponsiblesTo(aCollector)
	}
}

func assertValidTeam(aName string, aResponsibles []Responsible) {
	assertValidTeamName(aName)
	assertNotEmptyResponsible(aResponsibles)
	assertNotRepeatedResponsible(aResponsibles)
}

func assertValidTeamName(aName string) {
	nameWithoutSpaces := strings.Replace(aName, " ", "", -1)
	if len(nameWithoutSpaces) == 0 {
		panic(errors.New("team name can not be empty"))
	}
}

func assertNotRepeatedResponsible(aResponsibles []Responsible) {
	aResponsibleCollector := []Responsible{}
	for _, responsible := range aResponsibles {
		responsible.AddResponsiblesTo(&aResponsibleCollector)
	}

	if len(internal.RepeatedElements(aResponsibleCollector)) > 0 {
		panic(errors.New("team can not have duplicated responsible"))
	}
}

func assertNotEmptyResponsible(aResponsibles []Responsible) {
	if len(aResponsibles) == 0 {
		panic(errors.New("team must be composed of subteams or developers"))
	}
}

func (t Team) daysToCompleteTaskForEachResponsible(anEffort int) []time.Duration {
	daysToCompleteForEachResponsible := internal.Map(t.responsibles, func(aResponsible Responsible) time.Duration { return aResponsible.DaysToFinish(anEffort) })
	return daysToCompleteForEachResponsible
}
