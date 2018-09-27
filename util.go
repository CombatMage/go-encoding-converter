package main

import (
	"os"
)

func isFilePresent(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}
