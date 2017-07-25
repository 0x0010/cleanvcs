package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func DelVcsFiles(vd *VcsDirectory) {
	delFiles(vd.FilePath, false)
	delFiles(vd.DirPath, true)
}

func delFiles(files []string, isDir bool) {
	for _, element := range files {
		time.Sleep(50 * time.Millisecond)
		err := os.Remove(element)
		if nil != err {
			log.Fatal(err)
		}
		if isDir {
			fmt.Println("RM D ", element)
		} else {
			fmt.Println("RM F ", element)
		}
	}
}
