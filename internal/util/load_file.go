package util

import (
	"log"
	"os"
	"path/filepath"
)

// ProcessFilesDirectory...
func ProcessFilesDirectory(filesDir string) string {
	if filepath.IsAbs(filesDir) {
		return filesDir
	}

	//TODO: Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %v", err)
		return filesDir
	}

	//TODO: Check if we're in a subdirectory (like tests)
	if filepath.Base(wd) == "tests" || filepath.Base(wd) == "controller" {
		//TODO: Go up directories until we find the project root
		for {
			parent := filepath.Dir(wd)
			if parent == wd {
				break
			}
			//TODO: Check if go.mod exists (indicates project root)
			if _, err := os.Stat(filepath.Join(parent, "go.mod")); err == nil {
				wd = parent
				break
			}
			wd = parent
		}
	}

	return filepath.Join(wd, filesDir)
}
