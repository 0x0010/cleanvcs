package utils

import (
	"path/filepath"
	"os"
	"log"
)

// Get current execute path
func CurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
