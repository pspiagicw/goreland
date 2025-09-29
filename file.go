package goreland

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/go-homedir"
	cp "github.com/otiai10/copy"
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

func CopyIgnoreGit(src string, dest string, ignore []string) {
	err := cp.Copy(src, dest, cp.Options{
		Skip: generateSkipFunc(ignore),
	})
	if err != nil {
		LogError("Error copying %s: %v", src, err)
	}
}

func generateSkipFunc(ignore []string) func(srcinfo fs.FileInfo, src string, dest string) (bool, error) {
	return func(srcinfo fs.FileInfo, src string, dest string) (bool, error) {
		if srcinfo.IsDir() && filepath.Base(src) == ".git" {
			return true, nil
		}
		for _, pattern := range ignore {
			match, err := filepath.Match(pattern, src)
			if err != nil {
				return false, err
			}
			if match {
				return true, nil
			}
		}
		return false, nil
	}
}
