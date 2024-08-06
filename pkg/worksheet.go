package pkg

import (
	"Project/internal"
	"time"
)

type WorkSheet struct {
	project Project
}

func NewWorkSheet(project Project) *WorkSheet {
	ws := new(WorkSheet)
	ws.project = project
	return ws
}

func (ws WorkSheet) Overassignments() map[*Developer][]map[time.Time]int {
	workingDatesByDeveloper := map[*Developer][]time.Time{}
	ws.project.AddWorkingDatesByDeveloperTo(workingDatesByDeveloper)

	overassignmentsByDeveloper := map[*Developer][]map[time.Time]int{}
	for developer := range workingDatesByDeveloper {
		overassignmentsByDeveloper[developer] = internal.CountOccurrences(workingDatesByDeveloper[developer])
	}

	return overassignmentsByDeveloper
}
