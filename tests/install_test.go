package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/pspiagicw/goreland"
)

func TestInstall(t *testing.T) {
	destdir, err := os.MkdirTemp("", "")

	if err != nil {
		t.Fatalf("Couldn't create directory: %v", err)
	}

	file, err := os.CreateTemp("/tmp", "")

	defer file.Close()
	file.WriteString(fmt.Sprint(rand.Int()))

	options := &goreland.InstallFileOptions{
		CreateDir:   true,
		Permissions: 755,
	}
	err = goreland.InstallFile(file.Name(), filepath.Join(destdir, filepath.Base(file.Name())), options)
	if err != nil {
		t.Fatalf("Error Installing File: %v", err)
	}

	_, err = os.Stat(filepath.Join(destdir, filepath.Base(file.Name())))
	if err != nil {
		t.Fatalf("Error with destfile: %v", err)
	}

}
