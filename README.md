版本控制文件清理工具，**目前支持subversion和git**。使用Go语言编写。
### 使用说明
cleanvcs扫描当前目录的版本控制文件，所以为了保证cleanvcs可在任意目录执行，请确认cleanvcs已加入系统Path。

### 源码编译
````shell
go get github.com/0x0010/cleanvcs
cd $GOPATH/src/github.com/0x0010/cleanvcs
go install
````
使用源码编译获得的可执行文件，已默认加入$GOPATH/bin下。
