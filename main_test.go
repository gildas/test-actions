package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAdd(t *testing.T) {
	assert.Equal(t, 4, add(2, 2))
}
