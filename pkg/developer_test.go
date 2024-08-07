package pkg

import (
	"Project/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeveloperTestSuite struct {
	suite.Suite
}

func TestDeveloperTestSuite(t *testing.T) {
	suite.Run(t, new(DeveloperTestSuite))
}

func (suite *DeveloperTestSuite) Test01DeveloperNameCanNotBeEmpty() {
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperNameErrorMessage, func() {
		NewDeveloper("", 8, 60)
	})
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperNameErrorMessage, func() {
		NewDeveloper(" ", 8, 60)
	})
}

func (suite DeveloperTestSuite) Test02DeveloperDedicationMustBePositive() {
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperDedicationErrorMessage, func() {
		NewDeveloper("Dan Ingalls", 0, 60)
	})
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperDedicationErrorMessage, func() {
		NewDeveloper("Dan Ingalls", -3, 60)
	})
}

func (suite *DeveloperTestSuite) Test03DeveloperRateMustBePositive() {
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperRateErrorMessage, func() {
		NewDeveloper("Dan Ingalls", 8, 0)
	})
	assert.PanicsWithError(suite.T(), internal.InvalidDeveloperRateErrorMessage, func() {
		NewDeveloper("Dan Ingalls", 8, -20)
	})
}
