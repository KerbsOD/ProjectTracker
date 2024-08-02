package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Developers
var danIngalls = NewDeveloper("Dan Ingalls", 8, 60)
var alanKay = NewDeveloper("Alan Kay", 6, 80)

// Teams
var danTeam = NewTeam("Dan team", []Responsible{danIngalls})
var parcMobTeam = NewTeam("Parc mob team", []Responsible{danIngalls, alanKay})

// Dates
var julyFirst = time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)
var julyTheSecond = time.Date(2024, time.July, 2, 0, 0, 0, 0, time.UTC)
var julyTheThird = time.Date(2024, time.July, 3, 0, 0, 0, 0, time.UTC)

func Test01ConcreteTaskFinishesInADayIfDeveloperDedicationsIsTaskEffort(t *testing.T) {
	task := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	assert.Equal(t, julyFirst, task.FinishDate())
}

func Test02ConcreteTaskDoesNotFinishesInADayIfDeveloperDedicationIsLessThanTaskEffort(t *testing.T) {
	task := NewConcreteTask("SS A", alanKay, julyFirst, 8, []Task{})
	assert.Equal(t, julyTheSecond, task.FinishDate())
}

func Test03ConcreteTaskStartsOnAfterDependentsFinish(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	task2 := NewConcreteTask("SS B", danIngalls, julyFirst, 8, []Task{task1})
	assert.Equal(t, julyTheSecond, task2.FinishDate())
}

func Test04ConcreteTaskDoesNotStartBeforeDesiredStartingDate(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	task2 := NewConcreteTask("SS B", alanKay, julyTheThird, 8, []Task{task1})
	assert.Equal(t, julyTheThird, task2.StartDate())
}

func Test05ConcreteTaskDoesNotStartTheSameDayItsDependentsEnd(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	task2 := NewConcreteTask("SS B", danIngalls, julyFirst, 8, []Task{task1})
	assert.NotEqual(t, julyFirst, task2.StartDate())
}

func Test06ConcreteTaskStartsAfterGreatestFinishDateInDependents(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	task2 := NewConcreteTask("SS B", danIngalls, julyFirst, 16, []Task{})
	task3 := NewConcreteTask("SS C", danIngalls, julyFirst, 16, []Task{task1, task2})
	assert.Equal(t, julyTheThird, task3.StartDate())
}

func Test07ConcreteTaskTimeToFinishForTeamWithOneDeveloperIsDeveloperTimeToFinish(t *testing.T) {
	task := NewConcreteTask("SS A", danTeam, julyFirst, 16, []Task{})
	assert.Equal(t, julyTheSecond, task.FinishDate())
}

func Test08ConcreteTaskTimeToFinishDependsOnSlowestDeveloper(t *testing.T) {
	task := NewConcreteTask("SS A", parcMobTeam, julyFirst, 16, []Task{})
	assert.Equal(t, julyTheThird, task.FinishDate())
}

func Test09ProjectStartDateWithOneSubTaskIsSubtaskStartDate(t *testing.T) {
	task := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(t, julyFirst, project.StartDate())
}

func Test10ProjectStartDateIsSubtasksEarliestStartDate(t *testing.T) {
	taskSSA := NewConcreteTask("SS A", danIngalls, julyTheThird, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", parcMobTeam, julyTheSecond, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(t, julyTheSecond, project.StartDate())
}

func Test11ProjectFinishDateWithOneSubTaskIsSubtaskFinishDate(t *testing.T) {
	task := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(t, julyFirst, project.FinishDate())
}

func Test12ProjectFinishDateIsSubtasksLatestFinishDate(t *testing.T) {
	taskSSA := NewConcreteTask("SS A", danIngalls, julyFirst, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", parcMobTeam, julyFirst, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(t, julyTheThird, project.FinishDate())
}

func Test13ProjectWithoutOverAssignmentsReturnsEmptyCollection(t *testing.T) {

}
