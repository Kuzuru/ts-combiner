package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func Remove(folder string) error {
	dir := folder

	// Use filepath.Walk to iterate through all the files in the directory tree
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// If the file is a .ts file, delete it
		if filepath.Ext(path) == ".ts" {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
