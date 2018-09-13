package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	iconv "github.com/djimenez/iconv-go"
)

const fromEncoding = "windows-1252"
const toEncoding = "utf-8"

func main() {
	inputDir := "input"
	outputDir := "output"
	os.Mkdir(outputDir, 0666)

	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Printf("handling directory %s\n", path)
			err := os.Mkdir(filepath.Join(outputDir, path), 0666)
			if err != nil {
				fmt.Printf("error while creating directory %s\n", err)
			}
		}
		return nil
	})

	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Printf("handling file %s\n", path)

			err := saveFileWithEncoding(path, filepath.Join(outputDir, path), info.Mode())
			if err != nil {
				fmt.Printf("error while copying file %s\n", err)
			}
		}
		return nil
	})
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
