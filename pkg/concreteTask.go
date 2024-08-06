package pkg

import (
	"Project/internal"
	"time"
)

type ConcreteTask struct {
	name         string
	responsible  Responsible
	expectedDate time.Time
	effort       int
	dependents   []Task
}

func NewConcreteTask(name string, responsible Responsible, desiredDate time.Time, effort int, dependents []Task) *ConcreteTask {
	t := new(ConcreteTask)
	t.name = name
	t.responsible = responsible
	t.expectedDate = desiredDate
	t.effort = effort
	t.dependents = dependents
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

func (ct ConcreteTask) latestFinishDateOfSubtasks() time.Time {
	finishDates := internal.Map(ct.dependents, func(aTask Task) time.Time { return aTask.FinishDate() })
	latestFinishDate := internal.MaxDateInArray(finishDates)
	return latestFinishDate
}

func (ct ConcreteTask) AddWorkingDatesByDeveloperTo(aWorkingDatesArrayByDeveloper map[*Developer][]time.Time) {
	ct.responsible.AddDatesToDeveloper(aWorkingDatesArrayByDeveloper, ct.workingDates())
}

func (ct ConcreteTask) workingDates() []time.Time {
	return internal.DatesBetween(ct.StartDate(), ct.FinishDate())
}
