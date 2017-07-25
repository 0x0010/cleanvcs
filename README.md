版本控制文件清理工具，**目前支持subversion和git**。使用Go语言编写。
### 使用说明
cleanvcs扫描当前目录的版本控制文件，所以为了保证cleanvcs可在任意目录执行，请确认cleanvcs已加入系统Path。
使用命令行进入被版本管控的目录，执行cleanvcs命令。cleanvcs首先会扫描当前目录下的所有与版本管控有关的文件和目录。
````shell
scan result summary
files:       312
directory:   255
[Dangerous] delete? (y/n):
````
当扫描结果包含版本文件或者目录，会提示删除。
删除文件有50ms间隔，以便误删可及时终止。*使用ctrl+c可强制终止程序*。
### 源码编译
````shell
go get github.com/0x0010/cleanvcs
cd $GOPATH/src/github.com/0x0010/cleanvcs
go install
````
使用源码编译获得的可执行文件，已默认加入$GOPATH/bin下。
