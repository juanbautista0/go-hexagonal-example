package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDependencies(t *testing.T) {
	dependencies := InitDependencies()

	assert.NotNil(t, dependencies)
	assert.NotNil(t, dependencies.UserRepository)
	assert.NotNil(t, dependencies.RdsClient)
}
