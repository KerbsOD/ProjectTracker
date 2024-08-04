package pkg

import "time"

type WorkSheet struct {
	project Project
}

func NewWorkSheet(project Project) *WorkSheet {
	ws := new(WorkSheet)
	ws.project = project
	return ws
}

func (ws WorkSheet) Overassignments() map[*Developer][]time.Time {
	overassignments := make(map[*Developer][]time.Time)
	return overassignments
}
