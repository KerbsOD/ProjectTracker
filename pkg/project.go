package pkg

import (
	"Project/internal"
	"time"
)

type Project struct {
	name     string
	subtasks []Task
}

func NewProject(aName string, aSliceOfTasks []Task) *Project {
	p := new(Project)
	p.name = aName
	p.subtasks = aSliceOfTasks
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

func (p Project) AddWorkingDatesForEachDeveloper(aWorkingDatesSliceForEachDeveloper map[*Developer][]time.Time) {
	for _, subtask := range p.subtasks {
		subtask.AddWorkingDatesForEachDeveloper(aWorkingDatesSliceForEachDeveloper)
	}
}
