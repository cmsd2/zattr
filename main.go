package main

import (
	"archive/zip"

	"github.com/cmsd2/zattr/cmd"
)

func main() {
	cmd.Execute()
}

func openArchive(path string) (*zip.ReadCloser, error) {
	return zip.OpenReader(path)
}
