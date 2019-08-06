# gogin
基于gin自定义脚手架


1、Mac下编译Linux, Windows平台的64位可执行程序：

$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test.go

$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go


2、Linux下编译Mac, Windows平台的64位可执行程序：

$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build test.go

$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go



3、Windows下编译Mac, Linux平台的64位可执行程序：

$ SET CGO_ENABLED=0SET GOOS=darwin3 SET GOARCH=amd64 go build test.go

$ SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build test.go

注：如果编译web等工程项目，直接cd到工程目录下直接执行以上命令

GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows

GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm

Golang version 1.5以前版本在首次交叉编译时还需要配置交叉编译环境：

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash
