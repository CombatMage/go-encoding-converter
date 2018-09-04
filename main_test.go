package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileWithEncoding(t *testing.T) {
	// action
	result, err := loadFileWithEncoding("testdata/input-1252/test.txt")
	// verify
	assert.NoError(t, err)
	assert.Equal(t, "üä@", result)
}

func TestSaveFileWithEncoding(t *testing.T) {
	// action
	err := saveFileWithEncoding("testdata/input-1252/test.txt", "testdata/out-utf8/")
	// verify
	assert.NoError(t, err)
}
