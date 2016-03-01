package main

import (
	"testing"

	"github.com/gambol99/go-marathon"
	"github.com/stretchr/testify/assert"
)

func TestMinHealthyTasksWhenEverythingIsFine(t *testing.T) {
	check := MinHealthyTasks{
		DefaultFailThreshold:    0.5,
		DefaultWarningThreshold: 0.6,
	}
	app := marathon.Application{
		ID:           "/foo",
		Instances:    100,
		TasksHealthy: 100,
	}

	appCheck := check.Check(app)
	assert.Equal(t, Pass, appCheck.Result)
	assert.Equal(t, "min-healthy", appCheck.CheckName)
	assert.Equal(t, "/foo", appCheck.App)
	assert.Equal(t, "We now have 100 healthy out of total 100", appCheck.Message)
}

func TestMinHealthyTasksWhenWarningThresholdIsMet(t *testing.T) {
	check := MinHealthyTasks{
		DefaultFailThreshold:    0.5,
		DefaultWarningThreshold: 0.6,
	}
	app := marathon.Application{
		ID:           "/foo",
		Instances:    100,
		TasksHealthy: 59,
	}

	appCheck := check.Check(app)
	assert.Equal(t, Warning, appCheck.Result)
	assert.Equal(t, "min-healthy", appCheck.CheckName)
	assert.Equal(t, "/foo", appCheck.App)
	assert.Equal(t, "Only 59 are healthy out of total 100", appCheck.Message)
}

func TestMinHealthyTasksWhenFailThresholdIsMet(t *testing.T) {
	check := MinHealthyTasks{
		DefaultFailThreshold:    0.5,
		DefaultWarningThreshold: 0.6,
	}
	app := marathon.Application{
		ID:           "/foo",
		Instances:    100,
		TasksHealthy: 49,
	}

	appCheck := check.Check(app)
	assert.Equal(t, Fail, appCheck.Result)
	assert.Equal(t, "min-healthy", appCheck.CheckName)
	assert.Equal(t, "/foo", appCheck.App)
	assert.Equal(t, "Only 49 are healthy out of total 100", appCheck.Message)
}

func TestMinHealthyTasksWhenNoTasksAreRunning(t *testing.T) {
	check := MinHealthyTasks{
		DefaultFailThreshold:    0.5,
		DefaultWarningThreshold: 0.6,
	}
	app := marathon.Application{
		ID:           "/foo",
		Instances:    1,
		TasksHealthy: 0,
	}

	appCheck := check.Check(app)
	assert.Equal(t, Fail, appCheck.Result)
	assert.Equal(t, "min-healthy", appCheck.CheckName)
	assert.Equal(t, "/foo", appCheck.App)
	assert.Equal(t, "Only 0 are healthy out of total 1", appCheck.Message)
}