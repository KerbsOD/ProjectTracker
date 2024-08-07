package pkg

import (
	"Project/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TeamTestSuite struct {
	suite.Suite
	danIngalls  *Developer
	alanKay     *Developer
	danTeam     *Team
	parcMobTeam *Team
}

func TestTeamTestSuite(t *testing.T) {
	suite.Run(t, new(TeamTestSuite))
}

func (suite *TeamTestSuite) SetupTest() {
	suite.danIngalls = NewDeveloper("Dan Ingalls", 8, 60)
	suite.alanKay = NewDeveloper("Alan Kay", 6, 80)
	suite.danTeam = NewTeam("Dan team", []Responsible{suite.danIngalls})
	suite.parcMobTeam = NewTeam("Parc mob team", []Responsible{suite.danIngalls, suite.alanKay})
}

func (suite *TeamTestSuite) Test01TeamNameCanNotBeEmpty() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamNameErrorMessage, func() {
		NewTeam("", []Responsible{suite.danIngalls})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidTeamNameErrorMessage, func() {
		NewTeam(" ", []Responsible{suite.danIngalls})
	})
}

func (suite *TeamTestSuite) Test02TeamMustBeComposedOfSubteamsOrDevelopers() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Parc mob team", []Responsible{})
	})
}

func (suite *TeamTestSuite) Test03TeamCanNotHaveRepeatedSubteamsOrDevelopers() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Team Dynamite", []Responsible{suite.alanKay, suite.alanKay})
	})
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Team Super Cool", []Responsible{suite.parcMobTeam, suite.parcMobTeam})
	})
}

func (suite *TeamTestSuite) Test04TeamCanNotHaveRepeatedSubteamsOrDevelopersWithinSubteams() {
	assert.PanicsWithError(suite.T(), internal.InvalidTeamResponsibleErrorMessage, func() {
		NewTeam("Team Super Cool", []Responsible{suite.alanKay, suite.parcMobTeam})
	})
}
