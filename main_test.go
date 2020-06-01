package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAdd(t *testing.T) {
	assert.Equal(t, 4, add(2, 2))
}

func TestHasProperEnv(t *testing.T) {
	data, ok := os.LookupEnv("GOOGLE_PROJECT_ID")
	assert.True(t, ok)
	assert.True(t, len(data) > 0)
}
