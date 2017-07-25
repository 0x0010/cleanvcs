这是一个清理版本控制文件的工具，使用Golang编写。
### 使用说明
cleanvcs工作原理是扫描当前目录的版本控制文件，所以为了保证cleanvcs可以在任意目录执行，请确认cleanvcs之前确认已加入系统Path。

如果是使用源码编译的方式，可执行文件已默认加入$GOPATH/bin下。

### 源码编译
````shell
go get github.com/0x0010/cleanvcs
cd $GOPATH/src/github.com/0x0010/cleanvcs
go install
````
