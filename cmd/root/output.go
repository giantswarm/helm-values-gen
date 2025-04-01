package root

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/giantswarm/microerror"
)

func handleOutput(flag *flag, marshalled []byte) error {
	output := flag.output
	force := flag.force

	if output == "" {
		fmt.Print(string(marshalled))
		return nil
	}

	if fileExists(output) && !force {
		return microerror.Maskf(outputError, "File %s already exists. Use --force to overwrite.", output)
	}

	file, err := createFile(output)
	if err != nil {
		return microerror.Mask(err)
	}
	defer func() { _ = file.Close() }()

	err = writeToFile(file, marshalled)
	if err != nil {
		return microerror.Mask(err)
	}

	log.Printf("\nWrote default values to %s", output)
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createFile(path string) (*os.File, error) {
	path = filepath.Clean(path)
	file, err := os.Create(path)
	if err != nil {
		return nil, microerror.Maskf(outputError, "Error creating file %s:\n%s", path, err)
	}
	return file, nil
}

func writeToFile(file *os.File, content []byte) error {
	_, err := file.Write(content)
	if err != nil {
		return microerror.Maskf(outputError, "Error writing to file %s:\n%s", file.Name(), err)
	}
	return nil
}
