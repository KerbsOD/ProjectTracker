package pkg

import (
	"Project/internal"
	"time"
)

type Project struct {
	name     string
	subtasks []Task
}

func NewProject(name string, subtasks []Task) *Project {
	p := new(Project)
	p.name = name
	p.subtasks = subtasks
	return p
}

func (p Project) StartDate() time.Time {
	earliestStartDate := p.earliestStartDateOfSubtasks()
	return earliestStartDate
}

func (p Project) FinishDate() time.Time {
	latestFinishDate := p.latestFinishDateOfSubtasks()
	return latestFinishDate
}

func (p Project) Worksheet() *WorkSheet {
	return NewWorkSheet(p)
}

func (p Project) AddWorkingDatesByDeveloperTo(aWorkingDatesArrayByDeveloper map[*Developer][]time.Time) {
	for _, subtask := range p.subtasks {
		subtask.AddWorkingDatesByDeveloperTo(aWorkingDatesArrayByDeveloper)
	}
}

func (p Project) earliestStartDateOfSubtasks() time.Time {
	startDates := internal.Map(p.subtasks, func(aTask Task) time.Time { return aTask.StartDate() })
	earliestFinishDate := internal.MinDateInArray(startDates)
	return earliestFinishDate
}

func (p Project) latestFinishDateOfSubtasks() time.Time {
	finishDates := internal.Map(p.subtasks, func(aTask Task) time.Time { return aTask.FinishDate() })
	latestFinishDate := internal.MaxDateInArray(finishDates)
	return latestFinishDate
}
