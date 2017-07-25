package utils

import (
	_log "github.com/0x0010/cleanvcs/log"
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
				vd.AppendDir(path)
				_log.Infof("D: %s", strings.TrimPrefix(path, root))
			} else {
				vd.AppendFile(path)
				_log.Infof("F: %s", strings.TrimPrefix(path, root))
			}
		}
		return nil
	}
	err := filepath.Walk(root, walkFn)
	if err != nil {
		log.Fatalf("Walk path %s failed", root)
	}
	return vd
}

// check path contains vcs signature, current only supports subversion
func isVcsPath(path string) bool {
	if len(path) <= 0 {
		return false
	}
	return strings.Contains(path, ".svn")
}

type VcsDirectory struct {
	filePath []string
	dirPath  []string
}

func NewVcsDirectory() *VcsDirectory {
	return &VcsDirectory{filePath: nil, dirPath: nil}
}

func (vd *VcsDirectory) AppendFile(file string) {
	vd.filePath = append(vd.filePath, file)
}

func (vd *VcsDirectory) AppendDir(dir string) {
	vd.dirPath = append(vd.dirPath, dir)
}

func (vd *VcsDirectory) DirCount() int {
	return len(vd.dirPath)
}

func (vd *VcsDirectory) FileCount() int {
	return len(vd.filePath)
}
