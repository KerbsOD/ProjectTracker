package app

import (
	"Project/internal/generics"
	"time"
)

type WorkSheet struct {
	project                 Project
	workingDatesByDeveloper map[*Developer][]time.Time
}

func NewWorkSheet(aProject Project) *WorkSheet {
	ws := new(WorkSheet)
	ws.project = aProject
	ws.workingDatesByDeveloper = ws.calculateWorkingDatesByDeveloper()
	return ws
}

func (ws WorkSheet) Overassignments() map[*Developer][]time.Time {
	overassignedDatesForEachDeveloper := make(map[*Developer][]time.Time)
	for developer, workingDates := range ws.workingDatesByDeveloper {
		overassignedDatesForEachDeveloper[developer] = ws.overassignedDates(workingDates)
	}

	return overassignedDatesForEachDeveloper
}

func (ws WorkSheet) HasOverassignments() bool {
	for _, overassignedDates := range ws.Overassignments() {
		if len(overassignedDates) > 0 {
			return true
		}
	}
	return false
}

func (ws WorkSheet) TotalCost() int {
	totalCost := 0
	for developer, workingDates := range ws.workingDatesByDeveloper {
		aNumberOfDays := len(workingDates)
		totalCost = totalCost + developer.CostForWorking(aNumberOfDays)
	}
	return totalCost
}

func (ws WorkSheet) overassignedDates(aWorkingDates []time.Time) []time.Time {
	return generics.RepeatedElements(aWorkingDates)
}

func (ws WorkSheet) calculateWorkingDatesByDeveloper() map[*Developer][]time.Time {
	workingDatesByDeveloper := map[*Developer][]time.Time{}
	ws.project.AddWorkingDatesForEachDeveloper(workingDatesByDeveloper)
	return workingDatesByDeveloper
}
