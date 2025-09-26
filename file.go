package goreland

import (
	"errors"
	"io/fs"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/go-homedir"
)

func Confirm(message string, failedMessage string) {
	confirm := false
	prompt := &survey.Confirm{
		Message: message,
	}

	survey.AskOne(prompt, &confirm)
	if !confirm {
		LogFatal(failedMessage)
	}
}

func DoesExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)

}

func CreateIfNotExist(folder string) {
	if _, err := os.Stat(folder); errors.Is(err, fs.ErrNotExist) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			LogFatal("Error creating directory: %s", folder)
		}
	} else if err != nil {
		LogFatal("Error stating file: %s", err)
	}
}

func ExpandHome(filepath string) string {
	path, err := homedir.Expand(filepath)
	if err != nil {
		LogFatal("Unable to expand home variable: %v", err)
	}
	return path
}
