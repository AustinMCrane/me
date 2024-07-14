package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"
)

var (
	outputPath = "public"
	inputPath  = "src"
)

func main() {
	args := os.Args[1:]
	// input dir and output dir
	if len(args) == 2 {
		outputPath = args[0]
		inputPath = args[1]
	}

	if err := Exec(inputPath, outputPath); err != nil {
		panic(err)
	}
}

func Exec(inputPath string, outputPath string) error {
	// read input dir
	templateFile, err := os.Open(fmt.Sprintf("%s/index.html", inputPath))
	if err != nil {
		return err
	}
	defer templateFile.Close()

	jsonResumeFile, err := os.Open(fmt.Sprintf("%s/resume.json", inputPath))
	if err != nil {
		return err
	}
	defer jsonResumeFile.Close()

	// read template file
	templStr, err := io.ReadAll(templateFile)
	if err != nil {
		return err
	}

	templ, err := template.New("foo").Parse(string(templStr))
	if err != nil {
		return err
	}

	// read json resume file
	jsonResumeData, err := io.ReadAll(jsonResumeFile)
	if err != nil {
		return err
	}
	var resume JSONResume
	err = json.Unmarshal(jsonResumeData, &resume)
	if err != nil {
		return err
	}

	indexPublicPath := fmt.Sprintf("%s/index.html", outputPath)
	indexPublicFile, err := os.Create(indexPublicPath)
	if err != nil {
		return err
	}

	// execute template
	err = templ.Execute(indexPublicFile, resume)
	if err != nil {
		return err
	}

	// copy over json resume
	jsonResumePublicPath := fmt.Sprintf("%s/resume.json", outputPath)
	jsonResumePublicFile, err := os.Create(jsonResumePublicPath)
	if err != nil {
		return err
	}
	err = json.NewEncoder(jsonResumePublicFile).Encode(resume)
	if err != nil {
		return err
	}

	// parse json resume
	return nil
}
