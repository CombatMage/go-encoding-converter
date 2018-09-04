package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	iconv "github.com/djimenez/iconv-go"
)

func main() {
	outputDir := "tmp"
	os.Mkdir(outputDir, 0666)

	filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			os.Mkdir(filepath.Join(outputDir, path), 0666)
		} else {
			saveFileWithEncoding(path, filepath.Join(outputDir, path))
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
