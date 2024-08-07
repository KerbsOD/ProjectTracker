package pkg

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type WorksheetTestSuite struct {
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

func TestWorksheetTestSuite(t *testing.T) {
	suite.Run(t, new(WorksheetTestSuite))
}

func (suite *WorksheetTestSuite) SetupTest() {
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

func (suite *WorksheetTestSuite) Test01DeveloperWithoutOverAssignmentsReturnsEmptyCollection() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()

	overassignments := make(map[*Developer][]time.Time)
	overassignments[suite.danIngalls] = []time.Time{}

	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.danIngalls], overassignments[suite.danIngalls])
}

func (suite *WorksheetTestSuite) Test02DeveloperWithOverassignmentsReturnsArrayWithOverassignedDays() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	project := NewProject("modelo", []Task{taskSSA, taskSSB})
	worksheet := project.Worksheet()

	overassignments := make(map[*Developer][]time.Time)
	overassignments[suite.danIngalls] = []time.Time{suite.july1th}
	overassignments[suite.alanKay] = []time.Time{}

	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.danIngalls], overassignments[suite.danIngalls])
	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.alanKay], overassignments[suite.alanKay])
}

func (suite *WorksheetTestSuite) Test03DevelopersWithOverassignmentsReturnsArrayWithOverassignedDaysPerDeveloper() {
	SSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	SSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	SSC := NewConcreteTask("SS C", suite.alanKay, suite.july2th, 16, []Task{SSA, SSB})
	model := NewProject("Modelo", []Task{SSA, SSB, SSC})
	UI := NewConcreteTask("UI", suite.adeleGoldberg, suite.july2th, 6, []Task{model})
	systemERP := NewProject("Sistema ERP", []Task{model, UI})

	worksheet := systemERP.Worksheet()

	overassignments := make(map[*Developer][]time.Time)
	overassignments[suite.danIngalls] = []time.Time{suite.july1th}
	overassignments[suite.alanKay] = []time.Time{}
	overassignments[suite.adeleGoldberg] = []time.Time{}

	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.danIngalls], overassignments[suite.danIngalls])
	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.alanKay], overassignments[suite.alanKay])
	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.adeleGoldberg], overassignments[suite.adeleGoldberg])
}

func (suite *WorksheetTestSuite) Test04ProjectDoesNotHaveOverassignmentsIfDevelopersWorkInOneTaskForEachDate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()
	assert.False(suite.T(), worksheet.HasOverassignments())
}

func (suite *WorksheetTestSuite) Test05ProjectHavesOverassignmentsIfADeveloperWorksInMoreThanOneTaskInADate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	project := NewProject("modelo", []Task{taskSSA, taskSSB})
	worksheet := project.Worksheet()
	assert.True(suite.T(), worksheet.HasOverassignments())
}

func (suite *WorksheetTestSuite) Test06ProjectWithOneDeveloperTotalCostIsDeveloperNumberOfWorkingDaysTimesRate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()
	assert.Equal(suite.T(), worksheet.TotalCost(), (8 * 60 * 1))
	/*
		- Dan Ingalls: 8*hour/day * 60*dollar/hour * 1*day = 480*dollar
		-> Total: 960*dollar
	*/
}

func (suite *WorksheetTestSuite) Test07ProjectTotalCostIsTheSumOfTasksCost() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.alanKay, suite.july1th, 6, []Task{})
	project := NewProject("modelo", []Task{taskSSA, taskSSB})
	worksheet := project.Worksheet()

	assert.Equal(suite.T(), worksheet.TotalCost(), (8*60*1)+(6*80*1))
	/*
		- Dan Ingalls: 8*hour/day * 60*dollar/hour * 1*day = 480*dollar
		- Alan Kay: 6*hour/day * 80*dollar/hour * 1*day = 480*dollar
		-> Total: 960*dollar
	*/
}

func (suite *WorksheetTestSuite) Test08ProjectOverAssignmentsSumToTotalCost() {
	SSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	SSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	SSC := NewConcreteTask("SS C", suite.alanKay, suite.july2th, 16, []Task{SSA, SSB})
	model := NewProject("Modelo", []Task{SSA, SSB, SSC})
	UI := NewConcreteTask("UI", suite.adeleGoldberg, suite.july2th, 6, []Task{model})
	systemERP := NewProject("Sistema ERP", []Task{model, UI})

	worksheet := systemERP.Worksheet()

	assert.Equal(suite.T(), worksheet.TotalCost(), 5450)
	/*
		- Dan Ingalls: 8*hour/day * 60*dollar/hour * 4*day = 1440*dollar (Dan is overassigned on July1th, so he charges double on that day)
		- Alan Kay: 6*hour/day * 80*dollar/hour * 6*day = 2880*dollar
		- Adele Goldberg: 10*hour/day * 65*dollar/hour * 1*day = 650*dollar
		-> Total: 5450*dollar
	*/
}
