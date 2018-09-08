package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	iconv "github.com/djimenez/iconv-go"
)

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
			err := saveFileWithEncoding(path, filepath.Join(outputDir, path))
			if err != nil {
				fmt.Printf("error while copying file %s\n", err)
			}
		}
		return nil
	})
}

func saveFileWithEncoding(file string, output string) error {
	content, err := loadFileWithEncoding(file)
	if err != nil {
		return err
	}
	return writeFile(content, output)
}

func loadFileWithEncoding(file string) (string, error) {
	in, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer in.Close()
	reader, err := iconv.NewReader(in, "windows-1252", "utf-8")
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(reader)
	return string(s), err
}

func writeFile(content string, file string) error {
	return ioutil.WriteFile(file, []byte(content), 0666)
}
