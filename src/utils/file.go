package utils

import (
	"errors"
	"io/ioutil"
)

func FindFileInDirectory(targetDir string, fileName string) error {

	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == fileName {
				return nil
			}
		}
	}

	return errors.New("not found")
}
