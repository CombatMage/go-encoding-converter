package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	iconv "github.com/djimenez/iconv-go"
)

const fromEncoding = "windows-1252"
const toEncoding = "utf-8"

func main() {
	inputDir := "input"
	outputDir := "output"

	createDirectoryStructure(inputDir, outputDir)
	copyFilesToDestination(inputDir, outputDir)
}

func createDirectoryStructure(srcDirectory string, dstDirectory string) error {
	input, err := os.Stat(srcDirectory)
	if err != nil {
		fmt.Printf("Could not open directory dir: %s\n", srcDirectory)
		return err
	}

	if !isFilePresent(dstDirectory) {
		err := os.Mkdir(dstDirectory, input.Mode())
		if err != nil {
			fmt.Printf("Could not create output directory: %s\n", srcDirectory)
			return err
		}
	}

	var errOut error
	srcDirectory = filepath.ToSlash(srcDirectory)
	filepath.Walk(srcDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error while traversing directory: %s\n", err)
			errOut = err
			return err
		}

		path = filepath.ToSlash(path)

		if path == srcDirectory {
			return nil
		}

		if info.IsDir() {

			if strings.HasPrefix(path, srcDirectory) {
				path = strings.Replace(path, srcDirectory, "", 1)
			}

			fmt.Printf("handling directory %s\n", path)
			newDirectory := filepath.Join(dstDirectory, path)
			err := os.Mkdir(newDirectory, info.Mode())
			if err != nil {
				fmt.Printf("error while creating directory: %s\n", err)
				errOut = err
				return err
			}
		}
		return nil
	})
	return errOut
}

func copyFilesToDestination(srcDirectory string, dstDirectory string) error {

	var errOut error
	srcDirectory = filepath.ToSlash(srcDirectory)
	filepath.Walk(srcDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error while traversing directory: %s\n", err)
			errOut = err
			return err
		}

		path = filepath.ToSlash(path)

		if path == srcDirectory {
			return nil
		}

		if !info.IsDir() {
			if strings.HasPrefix(path, srcDirectory) {
				path = strings.Replace(path, srcDirectory, "", 1)
			}

			fmt.Printf("handling file %s\n", path)
			newFile := filepath.Join(dstDirectory, path)
			err := saveFileWithEncoding(path, newFile, info.Mode())
			if err != nil {
				fmt.Printf("error while copying file: %s\n", err)
				errOut = err
				return err
			}
		}
		return nil
	})
	return errOut
}

func saveFileWithEncoding(file string, output string, mode os.FileMode) error {
	content, err := loadFileWithEncoding(file, fromEncoding, toEncoding)
	if err != nil {
		return err
	}
	return writeFile(content, output, mode)
}

func loadFileWithEncoding(file string, encodingIn string, encodingOut string) (string, error) {
	in, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer in.Close()
	reader, err := iconv.NewReader(in, encodingIn, encodingOut)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(reader)
	return string(s), err
}

func writeFile(content string, file string, mode os.FileMode) error {
	return ioutil.WriteFile(file, []byte(content), mode)
}
