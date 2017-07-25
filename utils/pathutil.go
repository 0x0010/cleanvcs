package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Get current execute path
func CurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func SearchVcs(vd *VcsDirectory, root string) *VcsDirectory {
	walkFn := func(path string, info os.FileInfo, err error) error {
		if isVcsPath(path) {
			if info.IsDir() {
				vd.DirPath = append(vd.DirPath, path)
				fmt.Printf("D: %s\n", strings.TrimPrefix(path, root))
			} else {
				vd.FilePath = append(vd.FilePath, path)
				fmt.Printf("F: %s\n", strings.TrimPrefix(path, root))
			}
		}
		return nil
	}
	err := filepath.Walk(root, walkFn)
	if err != nil {
		log.Fatalf("Walk path %s failed\n", root)
	}
	return vd
}

// check path contains vcs signature, current only supports subversion
func isVcsPath(path string) bool {
	if len(path) <= 0 {
		return false
	}
	return strings.Contains(path, ".svn") ||
		strings.Contains(path, ".git")
}

type VcsDirectory struct {
	FilePath []string
	DirPath  []string
}

func NewVcsDirectory() *VcsDirectory {
	return &VcsDirectory{FilePath: nil, DirPath: nil}
}
