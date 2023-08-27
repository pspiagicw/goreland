package goreland

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type InstallFileOptions struct {
	CreateDir   bool
	Permissions int
}

func InstallFile(sourcefile string, dest string, opts *InstallFileOptions) error {
	if opts.CreateDir {
		_, err := os.Stat(dest)
		if os.IsNotExist(err) {
			err = os.MkdirAll(dest, 755)
			if err != nil {
				return err
			}
		}
	}

	filename := filepath.Base(sourcefile)
	destfile := filepath.Join(dest, filename)

	in, err := os.Open(sourcefile)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	out, err := os.Create(destfile)
	if err != nil {
		in.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	in.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %s", err)
	}

	err = os.Chmod(destfile, fs.FileMode(opts.Permissions))
	if err != nil {
		return fmt.Errorf("Chmod error: %s", err)
	}

	return nil
}
