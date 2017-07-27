package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// CurrentPath getter
func CurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// SearchVcs aims to search vcs files under root directory
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
	return strings.Contains(path, fmt.Sprintf("%s.svn%s", pathSeparator(), pathSeparator())) ||
		strings.Contains(path, fmt.Sprintf("%s.git%s", pathSeparator(), pathSeparator()))
}

//VcsDirectory designed to store vcs files
type VcsDirectory struct {
	FilePath []string
	DirPath  []string
}

// NewVcsDirectory is VcsDirectory constructor
// exported to other package
func NewVcsDirectory() *VcsDirectory {
	return &VcsDirectory{FilePath: nil, DirPath: nil}
}

func pathSeparator() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\\"
	default:
		return "/"
	}
}
