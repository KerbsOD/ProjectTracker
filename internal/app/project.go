package app

import (
	"Project/internal/errorMessage"
	"Project/internal/extensions"
	"Project/internal/generics"
	"errors"
	"time"
)

type Project struct {
	name     string
	subtasks []Task
}

func NewProject(aName string, aSliceOfTasks []Task) *Project {
	assertValidProject(aName, aSliceOfTasks)
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

func (p Project) AddWorkingDatesForEachDeveloper(aWorkingDatesSliceForEachDeveloper map[*Developer][]time.Time) {
	for _, subtask := range p.subtasks {
		subtask.AddWorkingDatesForEachDeveloper(aWorkingDatesSliceForEachDeveloper)
	}
}

func (p Project) Worksheet() *WorkSheet {
	return NewWorkSheet(p)
}

/*
	PRIVATE
*/

func assertValidProject(aName string, aSliceOfTasks []Task) {
	assertValidProjectName(aName)
	assertValidSubtasks(aSliceOfTasks)
}

func assertValidProjectName(aName string) {
	if generics.EmptyName(aName) {
		panic(errors.New(errorMessage.InvalidProjectNameErrorMessage))
	}
}

func assertValidSubtasks(aSliceOfTasks []Task) {
	if len(aSliceOfTasks) == 0 {
		panic(errors.New(errorMessage.InvalidProjectSubtasksErrorMessage))
	}
}

func (p Project) earliestStartDateOfSubtasks() time.Time {
	startDates := generics.Map(p.subtasks, func(aTask Task) time.Time { return aTask.StartDate() })
	return extensions.MinDateInArray(startDates)
}

func (p Project) latestFinishDateOfSubtasks() time.Time {
	finishDates := generics.Map(p.subtasks, func(aTask Task) time.Time { return aTask.FinishDate() })
	return extensions.MaxDateInArray(finishDates)
}
