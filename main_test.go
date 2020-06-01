package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	filename, ok := os.LookupEnv("GOOGLE_CREDENTIALS")
	assert.True(t, ok)
	payload, err := ioutil.ReadFile(filepath.Join(".", filename))
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(payload), "END PRIVATE KEY"))
	t.Logf("Filename=%s", filename)
	t.Logf("Length=%d", len(string(payload)))
	t.Logf("payload=%s", string(payload))
}
