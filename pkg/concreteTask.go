package pkg

import (
	"Project/internal"
	"errors"
	"strings"
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

func assertValidConcreteTask(aName string, anEffort int, aSliceOfDependentTasks []Task) {
	assertValidConcreteTaskName(aName)
	assertValidEffort(anEffort)
	assertValidDependents(aSliceOfDependentTasks)
}

func assertValidDependents(aSliceOfDependentTasks []Task) {
	if len(internal.RepeatedElements(aSliceOfDependentTasks)) > 0 {
		panic(errors.New("concrete task can not have direct repeated tasks"))
	}
}

func assertValidEffort(anEffort int) {
	if anEffort <= 0 {
		panic(errors.New("concrete task effort must be positive"))
	}
}

func assertValidConcreteTaskName(aName string) {
	nameWithoutSpaces := strings.Replace(aName, " ", "", -1)
	if len(nameWithoutSpaces) == 0 {
		panic(errors.New("concrete task name can not be empty"))
	}
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

func (ct ConcreteTask) latestFinishDateOfSubtasks() time.Time {
	finishDates := internal.Map(ct.dependents, func(aTask Task) time.Time { return aTask.FinishDate() })
	latestFinishDate := internal.MaxDateInArray(finishDates)
	return latestFinishDate
}

func (ct ConcreteTask) AddWorkingDatesForEachDeveloper(aWorkingDatesArrayForEachDeveloper map[*Developer][]time.Time) {
	taskWorkingDates := ct.workingDates()
	ct.responsible.AddWorkingDatesForEachDeveloper(taskWorkingDates, aWorkingDatesArrayForEachDeveloper)
}

func (ct ConcreteTask) workingDates() []time.Time {
	return internal.DatesBetween(ct.StartDate(), ct.FinishDate())
}
