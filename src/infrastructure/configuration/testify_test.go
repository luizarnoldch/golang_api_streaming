package configuration_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}

func (suite *ExampleSuite) TestTrue() {
	suite.T().Log("...Running TestTrue")
	suite.True(true)
}

func (suite *ExampleSuite) TestFalse() {
	suite.T().Log("...Running TestFalse")
	suite.False(false)
}

func (suite *ExampleSuite) SetupSuite() {
	suite.T().Log("SetupSuite")
}

func (suite *ExampleSuite) TearDownSuite() {
	suite.T().Log("TearDownSuite")
}

func (suite *ExampleSuite) SetupTest() {
	suite.T().Log("SetupTest")
}

func (suite *ExampleSuite) TearDownTest() {
	suite.T().Log("TearDownTest")
}

func (suite *ExampleSuite) BeforeTest(suiteName, testName string) {
	suite.T().Log("BeforeTest")
}

func (suite *ExampleSuite) AfterTest(suiteName, testName string) {
	suite.T().Log("AfterTest")
}
