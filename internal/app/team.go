package app

import (
	"Project/internal/errorMessage"
	"Project/internal/extensions"
	"Project/internal/generics"
	"errors"
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
	maxDaysToCompleteTaskAmongResponsibles := generics.MaximizeElementByComparer(daysToCompleteTaskForEachResponsible, extensions.GreaterDuration)
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

/*
	PRIVATE
*/

func assertValidTeam(aName string, aResponsibles []Responsible) {
	assertValidTeamName(aName)
	assertValidResponsible(aResponsibles)
}

func assertValidTeamName(aName string) {
	if generics.EmptyName(aName) {
		panic(errors.New(errorMessage.InvalidTeamNameErrorMessage))
	}
}

func assertValidResponsible(aResponsibles []Responsible) {
	assertNotEmptyResponsible(aResponsibles)
	assertNotRepeatedResponsible(aResponsibles)
}

func assertNotEmptyResponsible(aResponsibles []Responsible) {
	if len(aResponsibles) == 0 {
		panic(errors.New(errorMessage.InvalidTeamResponsibleErrorMessage))
	}
}

func assertNotRepeatedResponsible(aResponsibles []Responsible) {
	aResponsibleCollector := []Responsible{}
	for _, responsible := range aResponsibles {
		responsible.AddResponsiblesTo(&aResponsibleCollector)
	}

	if len(generics.RepeatedElements(aResponsibleCollector)) > 0 {
		panic(errors.New(errorMessage.InvalidTeamResponsibleErrorMessage))
	}
}

func (t Team) daysToCompleteTaskForEachResponsible(anEffort int) []time.Duration {
	daysToCompleteForEachResponsible := generics.Map(t.responsibles, func(aResponsible Responsible) time.Duration { return aResponsible.DaysToFinish(anEffort) })
	return daysToCompleteForEachResponsible
}
