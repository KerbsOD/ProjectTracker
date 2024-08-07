package pkg

import (
	"Project/internal"
	"errors"
	"time"
)

type ConcreteTask struct {
	name         string
	responsible  Responsible
	expectedDate time.Time
	effort       int
	dependents   []Task
}

func NewConcreteTask(aName string, aResponsible Responsible, aDesiredStartingDate time.Time, anEffort int, aSliceOfDependentTasks []Task) *ConcreteTask {
	assertValidConcreteTask(aName, anEffort, aSliceOfDependentTasks)
	t := new(ConcreteTask)
	t.name = aName
	t.responsible = aResponsible
	t.expectedDate = aDesiredStartingDate
	t.effort = anEffort
	t.dependents = aSliceOfDependentTasks
	return t
}

func (ct ConcreteTask) StartDate() time.Time {
	if len(ct.dependents) == 0 {
		return ct.expectedDate
	}
	latestFinishDate := ct.latestFinishDateOfSubtasks()
	startDate := internal.MaxDateBetween(ct.expectedDate, latestFinishDate.Add(internal.Day)) // We don't start a task the same day we finish a task so the finish date will be the next day
	return startDate
}

func (ct ConcreteTask) FinishDate() time.Time {
	daysOfWork := ct.responsible.DaysToFinish(ct.effort)
	finishDate := ct.StartDate().Add(daysOfWork - internal.Day) // We finish at the end of the day, not the next day.
	return finishDate
}

func (ct ConcreteTask) AddWorkingDatesForEachDeveloper(aWorkingDatesArrayForEachDeveloper map[*Developer][]time.Time) {
	taskWorkingDates := ct.workingDates()
	ct.responsible.AddWorkingDatesForEachDeveloper(taskWorkingDates, aWorkingDatesArrayForEachDeveloper)
}

/*
	PRIVATE
*/

func assertValidConcreteTask(aName string, anEffort int, aSliceOfDependentTasks []Task) {
	assertValidConcreteTaskName(aName)
	assertValidEffort(anEffort)
	assertValidDependents(aSliceOfDependentTasks)
}

func assertValidConcreteTaskName(aName string) {
	if internal.EmptyName(aName) {
		panic(errors.New(internal.InvalidConcreteTaskNameErrorMessage))
	}
}

func assertValidEffort(anEffort int) {
	if anEffort <= 0 {
		panic(errors.New(internal.InvalidConcreteTaskEffortErrorMessage))
	}
}

func assertValidDependents(aSliceOfDependentTasks []Task) {
	if len(internal.RepeatedElements(aSliceOfDependentTasks)) > 0 {
		panic(errors.New(internal.InvalidConcreteTaskDependentsErrorMessage))
	}
}

func (ct ConcreteTask) latestFinishDateOfSubtasks() time.Time {
	finishDates := internal.Map(ct.dependents, func(aTask Task) time.Time { return aTask.FinishDate() })
	latestFinishDate := internal.MaxDateInArray(finishDates)
	return latestFinishDate
}

func (ct ConcreteTask) workingDates() []time.Time {
	return internal.DatesBetween(ct.StartDate(), ct.FinishDate())
}
