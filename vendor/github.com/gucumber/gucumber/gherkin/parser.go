package gherkin

import (
	"github.com/cucumber/gherkin-go"
	"os"
)

func ParseFile(filename string) (*gherkin.GherkinDocument, error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return gherkin.ParseGherkinDocument(f)
}
