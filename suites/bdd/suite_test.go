package bdd_test

import (
	"testing"

	"github.com/cucumber/godog"
)

func InitializeScenario(sc *godog.ScenarioContext) {
    initTodoSteps(sc)
}

// run with: go test -run TestFeatures -v
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "todo-api",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}