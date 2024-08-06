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

func (ws WorkSheet) Overassignments() map[*Developer][]time.Time {
	workingDatesByDeveloper := map[*Developer][]time.Time{}
	ws.project.AddWorkingDatesByDeveloperTo(workingDatesByDeveloper)

	overassignmentsByDeveloper := map[*Developer][]time.Time{}
	for developer := range workingDatesByDeveloper {
		overassignmentsByDeveloper[developer] = internal.ArrayWithRepeatedElements(workingDatesByDeveloper[developer])
	}

	return overassignmentsByDeveloper
}
