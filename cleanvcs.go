package main

import (
	"github.com/0x0010/cleanvcs/log"
	"github.com/0x0010/cleanvcs/pad"
	"github.com/0x0010/cleanvcs/utils"
)

func main() {
	curpath := utils.CurrentPath()
	log.Info("current path: ", curpath)
	log.Info("scanning vcs files and directories...")

	vd := utils.NewVcsDirectory()
	vd = utils.SearchVcs(vd, curpath)

	log.Info()
	log.Info("scan result summary")
	log.Info(pad.Right("files:", 12, " "), vd.FileCount())
	log.Info(pad.Right("directory:", 12, " "), vd.DirCount())
}
