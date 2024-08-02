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

func (p *Project) StartDate() time.Time {
	earliestStartDate := p.earliestStartDateOfSubtasks()
	return earliestStartDate
}

func (p *Project) FinishDate() time.Time {
	latestFinishDate := p.latestFinishDateOfSubtasks()
	return latestFinishDate
}

func (p Project) earliestStartDateOfSubtasks() time.Time {
	startDates := internal.Map(p.subtasks, func(aTask Task) time.Time { return aTask.StartDate() })
	earliestFinishDate := internal.MaxElement(startDates, time.Time.Before)
	return earliestFinishDate
}

func (p Project) latestFinishDateOfSubtasks() time.Time {
	finishDates := internal.Map(p.subtasks, func(aTask Task) time.Time { return aTask.FinishDate() })
	latestFinishDate := internal.MaxElement(finishDates, time.Time.After)
	return latestFinishDate
}
