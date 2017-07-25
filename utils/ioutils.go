package utils

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const SLEEP_TIME = 50 * time.Millisecond

func DelVcsFiles(vd *VcsDirectory) {
	delFiles(vd.FilePath)
	delDir(vd.DirPath)
}

func delFiles(files []string) {
	for _, element := range files {
		// sleep 50ms for risk shutdown
		time.Sleep(SLEEP_TIME)

		err := os.Remove(element)
		if nil != err {
			log.Fatal(err)
		}
		fmt.Println("RM F ", element)
	}
}

// sort directory slice and delete from end to start
func delDir(dir []string) {

	if !sort.StringsAreSorted(dir) {
		sort.Strings(dir)
	}

	for idx := len(dir) - 1; idx >= 0; idx-- {
		// sleep 50ms for risk shutdown
		time.Sleep(SLEEP_TIME)
		dirPath := dir[idx]
		err := os.Remove(dirPath)
		if nil != err {
			log.Fatal(err)
		}
		fmt.Println("RM D ", dirPath)
	}
}
