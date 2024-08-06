package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Developers
var danIngalls = NewDeveloper("Dan Ingalls", 8, 60)
var alanKay = NewDeveloper("Alan Kay", 6, 80)
var adeleGoldberg = NewDeveloper("Adele Goldberg", 10, 65)

// Teams
var danTeam = NewTeam("Dan team", []Responsible{danIngalls})
var parcMobTeam = NewTeam("Parc mob team", []Responsible{danIngalls, alanKay})

// Dates
var july1th = time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)
var july2th = time.Date(2024, time.July, 2, 0, 0, 0, 0, time.UTC)
var july3th = time.Date(2024, time.July, 3, 0, 0, 0, 0, time.UTC)
var july4th = time.Date(2024, time.July, 4, 0, 0, 0, 0, time.UTC)
var july5th = time.Date(2024, time.July, 5, 0, 0, 0, 0, time.UTC)
var july6th = time.Date(2024, time.July, 6, 0, 0, 0, 0, time.UTC)
var july7th = time.Date(2024, time.July, 7, 0, 0, 0, 0, time.UTC)

func Test01ConcreteTaskFinishesInADayIfDeveloperDedicationsIsTaskEffort(t *testing.T) {
	task := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	assert.Equal(t, july1th, task.FinishDate())
}

func Test02ConcreteTaskDoesNotFinishesInADayIfDeveloperDedicationIsLessThanTaskEffort(t *testing.T) {
	task := NewConcreteTask("SS A", alanKay, july1th, 8, []Task{})
	assert.Equal(t, july2th, task.FinishDate())
}

func Test03ConcreteTaskStartsOnAfterDependentsFinish(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", danIngalls, july1th, 8, []Task{task1})
	assert.Equal(t, july2th, task2.FinishDate())
}

func Test04ConcreteTaskDoesNotStartBeforeDesiredStartingDate(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", alanKay, july3th, 8, []Task{task1})
	assert.Equal(t, july3th, task2.StartDate())
}

func Test05ConcreteTaskDoesNotStartTheSameDayItsDependentsEnd(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", danIngalls, july1th, 8, []Task{task1})
	assert.NotEqual(t, july1th, task2.StartDate())
}

func Test06ConcreteTaskStartsAfterGreatestFinishDateInDependents(t *testing.T) {
	task1 := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", danIngalls, july1th, 16, []Task{})
	task3 := NewConcreteTask("SS C", danIngalls, july1th, 16, []Task{task1, task2})
	assert.Equal(t, july3th, task3.StartDate())
}

func Test07ConcreteTaskTimeToFinishForTeamWithOneDeveloperIsDeveloperTimeToFinish(t *testing.T) {
	task := NewConcreteTask("SS A", danTeam, july1th, 16, []Task{})
	assert.Equal(t, july2th, task.FinishDate())
}

func Test08ConcreteTaskTimeToFinishDependsOnSlowestDeveloper(t *testing.T) {
	task := NewConcreteTask("SS A", parcMobTeam, july1th, 16, []Task{})
	assert.Equal(t, july3th, task.FinishDate())
}

func Test09ProjectStartDateWithOneSubTaskIsSubtaskStartDate(t *testing.T) {
	task := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(t, july1th, project.StartDate())
}

func Test10ProjectStartDateIsSubtasksEarliestStartDate(t *testing.T) {
	taskSSA := NewConcreteTask("SS A", danIngalls, july3th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", parcMobTeam, july2th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(t, july2th, project.StartDate())
}

func Test11ProjectFinishDateWithOneSubTaskIsSubtaskFinishDate(t *testing.T) {
	task := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(t, july1th, project.FinishDate())
}

func Test12ProjectFinishDateIsSubtasksLatestFinishDate(t *testing.T) {
	taskSSA := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", parcMobTeam, july1th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(t, july3th, project.FinishDate())
}

func Test13DeveloperWithoutOverAssignmentsReturnsEmptyCollection(t *testing.T) {
	taskSSA := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()

	overassignments := make(map[*Developer][]map[time.Time]int)
	overassignments[danIngalls] = []map[time.Time]int{
		map[time.Time]int{july1th: 1},
	}

	assert.ElementsMatch(t, worksheet.Overassignments()[danIngalls], overassignments[danIngalls])
}

func Test14DeveloperWithOverassignmentsReturnsArrayWithOverassignedDays(t *testing.T) {
	taskSSA := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", parcMobTeam, july1th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	worksheet := project.Worksheet()

	overassignments := make(map[*Developer][]map[time.Time]int)
	overassignments[danIngalls] = []map[time.Time]int{
		map[time.Time]int{july1th: 2},
		map[time.Time]int{july2th: 1},
		map[time.Time]int{july3th: 1},
	}
	overassignments[alanKay] = []map[time.Time]int{
		map[time.Time]int{july1th: 1},
		map[time.Time]int{july2th: 1},
		map[time.Time]int{july3th: 1},
	}

	assert.ElementsMatch(t, worksheet.Overassignments()[danIngalls], overassignments[danIngalls])
	assert.ElementsMatch(t, worksheet.Overassignments()[alanKay], overassignments[alanKay])
}

func Test15DevelopersWithOverassignmentsReturnsArrayWithOverassignedDaysPerDeveloper(t *testing.T) {
	SSA := NewConcreteTask("SS A", danIngalls, july1th, 8, []Task{})
	SSB := NewConcreteTask("SS B", parcMobTeam, july1th, 16, []Task{})
	SSC := NewConcreteTask("SS C", alanKay, july2th, 16, []Task{SSA, SSB})
	model := NewProject("Modelo", []Task{SSA, SSB, SSC})
	UI := NewConcreteTask("UI", adeleGoldberg, july2th, 6, []Task{model})
	systemERP := NewProject("Sistema ERP", []Task{model, UI})

	worksheet := systemERP.Worksheet()

	overassignments := make(map[*Developer][]map[time.Time]int)
	overassignments[danIngalls] = []map[time.Time]int{
		map[time.Time]int{july1th: 2},
		map[time.Time]int{july2th: 1},
		map[time.Time]int{july3th: 1},
	}
	overassignments[alanKay] = []map[time.Time]int{
		map[time.Time]int{july1th: 1},
		map[time.Time]int{july2th: 1},
		map[time.Time]int{july3th: 1},
		map[time.Time]int{july4th: 1},
		map[time.Time]int{july5th: 1},
		map[time.Time]int{july6th: 1},
	}
	overassignments[adeleGoldberg] = []map[time.Time]int{
		map[time.Time]int{july7th: 1},
	}

	assert.ElementsMatch(t, worksheet.Overassignments()[danIngalls], overassignments[danIngalls])
	assert.ElementsMatch(t, worksheet.Overassignments()[alanKay], overassignments[alanKay])
	assert.ElementsMatch(t, worksheet.Overassignments()[adeleGoldberg], overassignments[adeleGoldberg])
}
