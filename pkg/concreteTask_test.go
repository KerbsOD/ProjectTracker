package pkg

import (
	"Project/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ConcreteTaskTestSuite struct {
	suite.Suite
	danIngalls    *Developer
	alanKay       *Developer
	adeleGoldberg *Developer
	danTeam       *Team
	parcMobTeam   *Team
	july1th       time.Time
	july2th       time.Time
	july3th       time.Time
}

func TestConcreteTaskTestSuite(t *testing.T) {
	suite.Run(t, new(ConcreteTaskTestSuite))
}

func (suite *ConcreteTaskTestSuite) SetupTest() {
	// Developers
	suite.danIngalls = NewDeveloper("Dan Ingalls", 8, 60)
	suite.alanKay = NewDeveloper("Alan Kay", 6, 80)
	suite.adeleGoldberg = NewDeveloper("Adele Goldberg", 10, 65)

	// Teams
	suite.danTeam = NewTeam("Dan team", []Responsible{suite.danIngalls})
	suite.parcMobTeam = NewTeam("Parc mob team", []Responsible{suite.danIngalls, suite.alanKay})

	// Dates
	suite.july1th = time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)
	suite.july2th = time.Date(2024, time.July, 2, 0, 0, 0, 0, time.UTC)
	suite.july3th = time.Date(2024, time.July, 3, 0, 0, 0, 0, time.UTC)
}

func (suite *ConcreteTaskTestSuite) Test01ConcreteTaskNameCanNotBeEmpty() {
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskNameErrorMessage, func() {
		NewConcreteTask("", suite.danIngalls, suite.july1th, 8, []Task{})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskNameErrorMessage, func() {
		NewConcreteTask(" ", suite.danIngalls, suite.july1th, 8, []Task{})
	})
}

func (suite *ConcreteTaskTestSuite) Test02ConcreteTaskEffortMustBePositive() {
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskEffortErrorMessage, func() {
		NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 0, []Task{})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskEffortErrorMessage, func() {
		NewConcreteTask("SS B", suite.danIngalls, suite.july1th, -8, []Task{})
	})
}

func (suite *ConcreteTaskTestSuite) Test03ConcreteTaskCanNotHaveDirectRepeatedDependentsTasks() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskDependentsErrorMessage, func() {
		NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{taskSSA, taskSSA})
	})
	/*
		It may happen that task X depends on tasks Y and Z, and at the same time, task Z depends on task Y.
		Circular dependency is impossible because we can not add an object at its own construction (clash).
	*/
}

func (suite *ConcreteTaskTestSuite) Test04ConcreteTaskFinishesInADayIfDeveloperDedicationsIsTaskEffort() {
	task := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	assert.Equal(suite.T(), suite.july1th, task.FinishDate())
}

func (suite *ConcreteTaskTestSuite) Test05ConcreteTaskDoesNotFinishesInADayIfDeveloperDedicationIsLessThanTaskEffort() {
	task := NewConcreteTask("SS A", suite.alanKay, suite.july1th, 8, []Task{})
	assert.Equal(suite.T(), suite.july2th, task.FinishDate())
}

func (suite *ConcreteTaskTestSuite) Test06ConcreteTaskStartsOnAfterDependentsFinish() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.danIngalls, suite.july1th, 8, []Task{task1})
	assert.Equal(suite.T(), suite.july2th, task2.FinishDate())
}

func (suite *ConcreteTaskTestSuite) Test07ConcreteTaskDoesNotStartBeforeDesiredStartingDate() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.alanKay, suite.july3th, 8, []Task{task1})
	assert.Equal(suite.T(), suite.july3th, task2.StartDate())
}

func (suite *ConcreteTaskTestSuite) Test08ConcreteTaskDoesNotStartTheSameDayItsDependentsEnd() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.danIngalls, suite.july1th, 8, []Task{task1})
	assert.NotEqual(suite.T(), suite.july1th, task2.StartDate())
}

func (suite *ConcreteTaskTestSuite) Test09ConcreteTaskStartsAfterGreatestFinishDateInDependents() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.danIngalls, suite.july1th, 16, []Task{})
	task3 := NewConcreteTask("SS C", suite.danIngalls, suite.july1th, 16, []Task{task1, task2})
	assert.Equal(suite.T(), suite.july3th, task3.StartDate())
}

func (suite *ConcreteTaskTestSuite) Test10ConcreteTaskTimeToFinishForTeamWithOneDeveloperIsDeveloperTimeToFinish() {
	task := NewConcreteTask("SS A", suite.danTeam, suite.july1th, 16, []Task{})
	assert.Equal(suite.T(), suite.july2th, task.FinishDate())
}

func (suite *ConcreteTaskTestSuite) Test11ConcreteTaskTimeToFinishDependsOnSlowestDeveloper() {
	task := NewConcreteTask("SS A", suite.parcMobTeam, suite.july1th, 16, []Task{})
	assert.Equal(suite.T(), suite.july3th, task.FinishDate())
}
