package pkg

import (
	"Project/internal"
	"time"
)

type WorkSheet struct {
	project                            Project
	numberOfDailyTasksForEachDeveloper map[*Developer]map[time.Time]int
}

func NewWorkSheet(aProject Project) *WorkSheet {
	ws := new(WorkSheet)
	ws.project = aProject
	ws.numberOfDailyTasksForEachDeveloper = ws.numberOfTasksPerDateForEachDeveloper()
	return ws
}

func (ws WorkSheet) Overassignments() map[*Developer][]time.Time {
	overassignedDatesForEachDeveloper := make(map[*Developer][]time.Time)
	for developer, numberOfTasksForEachDate := range ws.numberOfDailyTasksForEachDeveloper {
		overassignedDatesForEachDeveloper[developer] = ws.overassignedDates(numberOfTasksForEachDate)
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

func (ws WorkSheet) overassignedDates(aNumberOfTasksForEachDate map[time.Time]int) []time.Time {
	overassignedDates := []time.Time{}
	for date, numberOfTasksInDate := range aNumberOfTasksForEachDate {
		if numberOfTasksInDate > 1 {
			overassignedDates = append(overassignedDates, date)
		}
	}
	return overassignedDates
}

func (ws WorkSheet) numberOfTasksPerDateForEachDeveloper() map[*Developer]map[time.Time]int {
	numberOfTasksForEachDateByDeveloper := map[*Developer]map[time.Time]int{}
	for developer, workingDates := range ws.workingDatesByDeveloper() {
		numberOfTasksForEachDateByDeveloper[developer] = internal.MapWithNumberOfOccurrencesForEachElement(workingDates)
	}
	return numberOfTasksForEachDateByDeveloper
}

func (ws WorkSheet) workingDatesByDeveloper() map[*Developer][]time.Time {
	workingDatesByDeveloper := map[*Developer][]time.Time{}
	ws.project.AddWorkingDatesForEachDeveloper(workingDatesByDeveloper)
	return workingDatesByDeveloper
}
