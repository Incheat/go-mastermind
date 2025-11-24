package bdd_test

import (
	"testing"

	"github.com/cucumber/godog"
)

// run with: go test -run TestFeatures -v
func TestTodoFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "test-todo-api",
		ScenarioInitializer: initTodoSteps,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features/bdd/todo.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestMastermindFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "test-mastermind-api",
		TestSuiteInitializer: initMastermindTestSuite,
		ScenarioInitializer: initMastermindScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features/bdd/mastermind.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}