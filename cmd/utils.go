package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var redBold = color.New(color.FgRed, color.Bold).PrintlnFunc()
var greenBold = color.New(color.FgGreen, color.Bold).PrintlnFunc()

func handleError(err error) error {
	if err != nil {
		redBold(err)
	}
	return err
}

func activeSession() (string, error) {
	matches, err := filepath.Glob("*.stash")
	if err != nil {
		return "", handleError(err)
	}
	if len(matches) == 0 {
		return "", handleError(fmt.Errorf("no active .stash file found: start a session first"))
	}

	stashFileName := matches[0]
	titleOnly := strings.TrimSuffix(stashFileName, ".stash")
	filename := fmt.Sprintf("%s.md", titleOnly)

	return filename, nil
}
