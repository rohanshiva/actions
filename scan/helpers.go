package main

import (
	"os"
	"path/filepath"
)

// fileExists returns a bool indicating if a certain file exists in a dir
func fileExists(dir, filename string) bool {
	info, err := os.Stat(filepath.Join(dir, filename))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// checkFiles returns a bool indicating if any filename exists in a dir
func checkFiles(filenames ...string) func(dir string) bool {
	return func(dir string) bool {
		for _, filename := range filenames {
			if fileExists(dir, filename) {
				return true
			}
		}
		return false
	}
}