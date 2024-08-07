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

func (suite *ProjectTestSuite) Test01ConcreteTaskFinishesInADayIfDeveloperDedicationsIsTaskEffort() {
	task := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	assert.Equal(suite.T(), suite.july1th, task.FinishDate())
}

func (suite *ProjectTestSuite) Test02ConcreteTaskDoesNotFinishesInADayIfDeveloperDedicationIsLessThanTaskEffort() {
	task := NewConcreteTask("SS A", suite.alanKay, suite.july1th, 8, []Task{})
	assert.Equal(suite.T(), suite.july2th, task.FinishDate())
}

func (suite *ProjectTestSuite) Test03ConcreteTaskStartsOnAfterDependentsFinish() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.danIngalls, suite.july1th, 8, []Task{task1})
	assert.Equal(suite.T(), suite.july2th, task2.FinishDate())
}

func (suite *ProjectTestSuite) Test04ConcreteTaskDoesNotStartBeforeDesiredStartingDate() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.alanKay, suite.july3th, 8, []Task{task1})
	assert.Equal(suite.T(), suite.july3th, task2.StartDate())
}

func (suite *ProjectTestSuite) Test05ConcreteTaskDoesNotStartTheSameDayItsDependentsEnd() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.danIngalls, suite.july1th, 8, []Task{task1})
	assert.NotEqual(suite.T(), suite.july1th, task2.StartDate())
}

func (suite *ProjectTestSuite) Test06ConcreteTaskStartsAfterGreatestFinishDateInDependents() {
	task1 := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	task2 := NewConcreteTask("SS B", suite.danIngalls, suite.july1th, 16, []Task{})
	task3 := NewConcreteTask("SS C", suite.danIngalls, suite.july1th, 16, []Task{task1, task2})
	assert.Equal(suite.T(), suite.july3th, task3.StartDate())
}

func (suite *ProjectTestSuite) Test07ConcreteTaskTimeToFinishForTeamWithOneDeveloperIsDeveloperTimeToFinish() {
	task := NewConcreteTask("SS A", suite.danTeam, suite.july1th, 16, []Task{})
	assert.Equal(suite.T(), suite.july2th, task.FinishDate())
}

func (suite *ProjectTestSuite) Test08ConcreteTaskTimeToFinishDependsOnSlowestDeveloper() {
	task := NewConcreteTask("SS A", suite.parcMobTeam, suite.july1th, 16, []Task{})
	assert.Equal(suite.T(), suite.july3th, task.FinishDate())
}

func (suite *ProjectTestSuite) Test09ProjectStartDateWithOneSubTaskIsSubtaskStartDate() {
	task := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(suite.T(), suite.july1th, project.StartDate())
}

func (suite *ProjectTestSuite) Test10ProjectStartDateIsSubtasksEarliestStartDate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july3th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july2th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(suite.T(), suite.july2th, project.StartDate())
}

func (suite *ProjectTestSuite) Test11ProjectFinishDateWithOneSubTaskIsSubtaskFinishDate() {
	task := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("UI", []Task{task})
	assert.Equal(suite.T(), suite.july1th, project.FinishDate())
}

func (suite *ProjectTestSuite) Test12ProjectFinishDateIsSubtasksLatestFinishDate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	project := NewProject("Modelo", []Task{taskSSA, taskSSB})
	assert.Equal(suite.T(), suite.july3th, project.FinishDate())
}

func (suite *ProjectTestSuite) Test13DeveloperWithoutOverAssignmentsReturnsEmptyCollection() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()

	overassignments := make(map[*Developer][]time.Time)
	overassignments[suite.danIngalls] = []time.Time{}

	assert.ElementsMatch(suite.T(), worksheet.Overassignments()[suite.danIngalls], overassignments[suite.danIngalls])
}

func (suite *ProjectTestSuite) Test14DeveloperWithOverassignmentsReturnsArrayWithOverassignedDays() {
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

func (suite *ProjectTestSuite) Test15DevelopersWithOverassignmentsReturnsArrayWithOverassignedDaysPerDeveloper() {
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

func (suite *ProjectTestSuite) Test16ProjectDoesNotHaveOverassignmentsIfDevelopersWorkInOneTaskForEachDate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()
	assert.False(suite.T(), worksheet.HasOverassignments())
}

func (suite *ProjectTestSuite) Test17ProjectHavesOverassignmentsIfADeveloperWorksInMoreThanOneTaskInADate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	taskSSB := NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{})
	project := NewProject("modelo", []Task{taskSSA, taskSSB})
	worksheet := project.Worksheet()
	assert.True(suite.T(), worksheet.HasOverassignments())
}

func (suite *ProjectTestSuite) Test18ProjectWithOneDeveloperTotalCostIsDeveloperNumberOfWorkingDaysTimesRate() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	project := NewProject("Modelo", []Task{taskSSA})
	worksheet := project.Worksheet()
	assert.Equal(suite.T(), worksheet.TotalCost(), (8 * 60 * 1))
	/*
		- Dan Ingalls: 8*hour/day * 60*dollar/hour * 1*day = 480*dollar
		-> Total: 960*dollar
	*/
}

func (suite *ProjectTestSuite) Test19ProjectTotalCostIsTheSumOfTasksCost() {
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

func (suite *ProjectTestSuite) Test20ProjectOverAssignmentsSumToTotalCost() {
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

func (suite *ProjectTestSuite) Test21DeveloperNameCanNotBeEmpty() {
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperNameErrorMessage, func() {
		NewDeveloper("", 8, 60)
	})
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperNameErrorMessage, func() {
		NewDeveloper(" ", 8, 60)
	})
}

func (suite *ProjectTestSuite) Test22DeveloperDedicationMustBePositive() {
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperDedicationErrorMessage, func() {
		NewDeveloper("Dan Ingalls", 0, 60)
	})
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperDedicationErrorMessage, func() {
		NewDeveloper("Dan Ingalls", -3, 60)
	})
}

func (suite *ProjectTestSuite) Test23DeveloperRateMustBePositive() {
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperRateErrorMessage, func() {
		NewDeveloper("Dan Ingalls", 8, 0)
	})
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperRateErrorMessage, func() {
		NewDeveloper("Dan Ingalls", 8, -20)
	})
}

func (suite *ProjectTestSuite) Test24TeamNameCanNotBeEmpty() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamNameErrorMessage, func() {
		NewTeam("", []Responsible{suite.danIngalls})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidTeamNameErrorMessage, func() {
		NewTeam(" ", []Responsible{suite.danIngalls})
	})
}

func (suite *ProjectTestSuite) Test25TeamMustBeComposedOfSubteamsOrDevelopers() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Parc mob team", []Responsible{})
	})
}

func (suite *ProjectTestSuite) Test26TeamCanNotHaveRepeatedSubteamsOrDevelopers() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Team Dynamite", []Responsible{suite.alanKay, suite.alanKay})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Team Super Cool", []Responsible{suite.parcMobTeam, suite.parcMobTeam})
	})
}

func (suite *ProjectTestSuite) Test27TeamCanNotHaveRepeatedSubteamsOrDevelopersWithinSubteams() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Team Super Cool", []Responsible{suite.alanKay, suite.parcMobTeam})
	})
}

func (suite *ProjectTestSuite) Test28ConcreteTaskNameCanNotBeEmpty() {
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskNameErrorMessage, func() {
		NewConcreteTask("", suite.danIngalls, suite.july1th, 8, []Task{})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskNameErrorMessage, func() {
		NewConcreteTask(" ", suite.danIngalls, suite.july1th, 8, []Task{})
	})
}

func (suite *ProjectTestSuite) Test29ConcreteTaskEffortMustBePositive() {
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskEffortErrorMessage, func() {
		NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 0, []Task{})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskEffortErrorMessage, func() {
		NewConcreteTask("SS B", suite.danIngalls, suite.july1th, -8, []Task{})
	})
}

func (suite *ProjectTestSuite) Test30ConcreteTaskCanNotHaveDirectRepeatedDependentsTasks() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	assert.PanicsWithError(suite.T(), internal.InvalidConcreteTaskDependentsErrorMessage, func() {
		NewConcreteTask("SS B", suite.parcMobTeam, suite.july1th, 16, []Task{taskSSA, taskSSA})
	})
	/*
		It may happen that task X depends on tasks Y and Z, and at the same time, task Z depends on task Y.
		Circular dependency is impossible because we can not add an object at its own construction (clash).
	*/
}

func (suite *ProjectTestSuite) Test31ProjectNameCanNotBeEmpty() {
	taskSSA := NewConcreteTask("SS A", suite.danIngalls, suite.july1th, 8, []Task{})
	assert.PanicsWithError(suite.T(), internal.InvalidProjectNameErrorMessage, func() {
		NewProject("", []Task{taskSSA})
	})
}

func (suite *ProjectTestSuite) Test32ProjectCanNotHaveEmptySubtasks() {
	assert.PanicsWithError(suite.T(), internal.InvalidProjectSubtasksErrorMessage, func() {
		NewProject("UI", []Task{})
	})
}
