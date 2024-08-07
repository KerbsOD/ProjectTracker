package pkg

import (
	"Project/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ProjectTestSuite struct {
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

func TestProjectTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectTestSuite))
}

func (suite *ProjectTestSuite) SetupTest() {
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

func (suite *ProjectTestSuite) Test01ProjectNameCanNotBeEmpty() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	assert.PanicsWithError(suite.T(), internal.InvalidProjectNameErrorMessage, func() {
		NewProject("", []Task{taskSSA})
	})
}

func (suite *ProjectTestSuite) Test02ProjectCanNotHaveEmptySubtasks() {
	assert.PanicsWithError(suite.T(), internal.InvalidProjectSubtasksErrorMessage, func() {
		NewProject("UI", []Task{})
	})
}

func (suite *ProjectTestSuite) Test03ProjectStartDateWithOneSubTaskIsSubtaskStartDate() {
	task := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(suite.T(), suite.july1th, project.StartDate())
}

func (suite *ProjectTestSuite) Test04ProjectStartDateIsSubtasksEarliestStartDate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july3th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july2th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(suite.T(), suite.july2th, project.StartDate())
}

func (suite *ProjectTestSuite) Test05ProjectFinishDateWithOneSubTaskIsSubtaskFinishDate() {
	task := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(suite.T(), suite.july1th, project.FinishDate())
}

func (suite *ProjectTestSuite) Test06ProjectFinishDateIsSubtasksLatestFinishDate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(suite.T(), suite.july3th, project.FinishDate())
}
