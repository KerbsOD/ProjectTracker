package pkg

import (
	"Project/internal"
	"time"
)

type WorkSheet struct {
	project Project
}

func NewWorkSheet(aProject Project) *WorkSheet {
	ws := new(WorkSheet)
	ws.project = aProject
	return ws
}

func (ws WorkSheet) Overassignments() map[*Developer][]map[time.Time]int {
	return ws.numberOfWorkingDatesByDeveloper()
}

func (ws WorkSheet) numberOfWorkingDatesByDeveloper() map[*Developer][]map[time.Time]int {
	numberOfTasksForEachDateByDeveloper := map[*Developer][]map[time.Time]int{}
	for developer, workingDates := range ws.workingDatesByDeveloper() {
		numberOfTasksForEachDateByDeveloper[developer] = internal.SliceWithNumberOfOccurrencesForEachElement(workingDates)
	}
	return numberOfTasksForEachDateByDeveloper
}

func (ws WorkSheet) workingDatesByDeveloper() map[*Developer][]time.Time {
	workingDatesByDeveloper := map[*Developer][]time.Time{}
	ws.project.AddWorkingDatesForEachDeveloper(workingDatesByDeveloper)
	return workingDatesByDeveloper
}
