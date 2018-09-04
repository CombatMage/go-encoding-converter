package main

import (
	"io/ioutil"
	"os"

	iconv "github.com/djimenez/iconv-go"
)

func main() {

}

func saveFileWithEncoding(file string, outputDir string) error {
	content, err := loadFileWithEncoding(file)
	if err != nil {
		return err
	}
	output := outputDir + file
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
