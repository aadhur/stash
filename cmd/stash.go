package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func start(title string, restart bool) (string, error) {
	filename := title + ".md"
	content := "# " + title + "\n\n"
	stashname := title + ".stash"

	// Check for active session first
	matches, err := filepath.Glob("*.stash")
	if err != nil {
		return "", handleError(err)
	}
	if len(matches) != 0 {
		return "", handleError(fmt.Errorf("Active session already running: end it first"))
	}

	if restart {
		// Check if .md file exists before creating .stash
		mdMatches, err := filepath.Glob(filename)
		if err != nil {
			return "", handleError(err)
		}
		if len(mdMatches) == 0 {
			return "", handleError(fmt.Errorf("no file named %s exists", filename))
		}

		f, err := os.OpenFile(stashname, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", handleError(err)
		}
		defer f.Close()
		_, err = f.WriteString(filename + "\n")
		if err != nil {
			return "", handleError(err)
		}

	} else {
		// New session — create .md and .stash
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			return "", handleError(err)
		}

		f, err := os.OpenFile(stashname, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", handleError(err)
		}
		defer f.Close()
		_, err = f.WriteString(filename + "\n")
		if err != nil {
			return "", handleError(err)
		}
	}

	return filename, nil
}

func log(command string, addtitle string) (string, error) {

	filename, err := activeSession()
	if err != nil {
		return "", err
	}

	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "", handleError(fmt.Errorf("empty command"))
	}

	var buf bytes.Buffer
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &buf)

	err = cmd.Run()
	if err != nil {
		return "", handleError(err)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "", handleError(err)
	}
	defer file.Close()

	if addtitle != "" {
		_, err = file.WriteString(fmt.Sprintf("\n## %s\n", addtitle))
		if err != nil {
			return "", handleError(err)
		}
	}

	entry := fmt.Sprintf(
		"\n### Command\n```bash\n%s\n```\n\n### Output\n```output\n%s\n```\n",
		command,
		buf.String(),
	)

	_, err = file.WriteString(entry)
	if err != nil {
		return "", handleError(err)
	}

	return "", nil
}

func comment(comment string, status string) (string, error) {

	filename, err := activeSession()
	if err != nil {
		return "", err
	}

	// Open markdown file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "", handleError(err)
	}

	defer file.Close()

	var entry string

	// Command title mode
	if status != "" {

		entry = fmt.Sprintf(
			"---\n\n## %s\n%s\n",
			status,
			comment,
		)

	} else {

		entry = fmt.Sprintf(
			"\n%s\n",
			comment,
		)
	}

	_, err = file.WriteString(entry)
	if err != nil {
		return "", err
	}

	return "", nil

}

func end() (string, error) {
	matches, err := filepath.Glob("*.stash")
	if err != nil {
		return "", handleError(err)
	}

	if len(matches) == 0 {
		return "", handleError(fmt.Errorf("no active .stash file found: Start a session first"))
	}

	stashFileName := matches[0]

	// 2. Fix: Use standard assignment (=) because 'err' is already declared above
	err = os.Remove(stashFileName)
	if err != nil {
		return "", handleError(err)
	}

	// 3. Fix: Return a clean empty string and nil error on success
	return "Stash session ended successfully.", nil
}
