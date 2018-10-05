package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileWithEncoding(t *testing.T) {
	// action
	result, err := loadFileWithEncoding("testdata/input-1252/test.txt", "windows-1252", "utf-8")
	// verify
	assert.NoError(t, err)
	assert.Equal(t, "üä@", result)
}

func TestSaveFileWithEncoding(t *testing.T) {
	// arrange
	os.MkdirAll("testdata/out-utf8", 0777)
	os.Remove("testdata/out-utf8/test.txt")
	// action
	err := saveFileWithEncoding("testdata/input-1252/test.txt", "testdata/out-utf8/test.txt", 0777)
	// verify
	assert.NoError(t, err)
	result, _ := loadFileWithEncoding("testdata/out-utf8/test.txt", "utf-8", "utf-8")
	assert.Equal(t, "üä@", result)
}

func TestCreateDirectoryStructure(t *testing.T) {
	// arrange
	os.RemoveAll("testdata/output-directories")
	// action
	err := createDirectoryStructure("testdata/input-directories", "testdata/output-directories")
	// verify
	assert.NoError(t, err)
	assert.True(t, isFilePresent("testdata/output-directories/a"))
	assert.True(t, isFilePresent("testdata/output-directories/a/a1"))
	assert.True(t, isFilePresent("testdata/output-directories/a/a2"))
	assert.True(t, isFilePresent("testdata/output-directories/b"))
	assert.True(t, isFilePresent("testdata/output-directories/c"))
}

func TestCopyFilesToDestination(t *testing.T) {
	// arrange
	os.RemoveAll("testdata/output-directories")
	createDirectoryStructure("testdata/input-directories", "testdata/output-directories")
	// action
	err := copyFilesToDestination("testdata/input-directories", "testdata/output-directories")
	// verify
	assert.NoError(t, err)
	assert.True(t, isFilePresent("testdata/output-directories/b/b.txt"))
	assert.True(t, isFilePresent("testdata/output-directories/c/c.txt"))
}