package main

import (
	"bufio"
	"fmt"
	"github.com/0x0010/cleanvcs/pad"
	"github.com/0x0010/cleanvcs/utils"
	"os"
	"runtime"
	"strings"
)

var curr_path = utils.CurrentPath()

func main() {

	// log greeting to user
	logGreeting()

	// scan version control system files
	vd := utils.NewVcsDirectory()
	vd = utils.SearchVcs(vd, curr_path)

	// log scan result summary
	logSummary(vd)

	if len(vd.DirPath) > 0 || len(vd.FilePath) > 0 {
		fmt.Println("[Dangerous] delete? (y/n):")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if strings.EqualFold(strings.TrimSpace(text), "y") {
			utils.DelVcsFiles(vd)
		}
	} else {
		if strings.EqualFold(runtime.GOOS, "windows") {
			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}

func logGreeting() {
	fmt.Println("current path: ", curr_path)
	fmt.Println("scanning vcs files and directories...")
}

func logSummary(vd *utils.VcsDirectory) {
	fmt.Println()
	fmt.Println("scan result summary")
	fmt.Println(pad.Right("files:", 12, " "), len(vd.FilePath))
	fmt.Println(pad.Right("directory:", 12, " "), len(vd.DirPath))
}
