package goversion

import (
	"flag"
	"fmt"
	"os"
)

//编译相关的信息
var (
	BuildVersion  string // 编译版本
	BuildTime     string // 编译时间
	BuildName     string // 编译程序名称
	GitCommitID   string // git 最后的提交 commit
	GitBranch     string // git branch
	GolangVersion string // golang的版本
	HostName      string // 编译机器
)

var showV = flag.Bool("v", false, "show the app version")
var showVersion = flag.Bool("version", false, "show the app version")

//ShowAndExit 显示并退出程序
func ShowAndExit() {
	showVersionAndExit()
}

//显示版本并退出
func showVersionAndExit() {
	if *showV || *showVersion {
		fmt.Printf("************************************************************\n")
		fmt.Printf("* build name:      %s\n", BuildName)
		fmt.Printf("* build version:   %s\n", BuildVersion)
		fmt.Printf("* build time:      %s\n", BuildTime)
		fmt.Printf("* git commit:      %s\n", GitCommitID)
		fmt.Printf("* git branch:      %s\n", GitBranch)
		fmt.Printf("* golang version:  %s\n", GolangVersion)
		fmt.Printf("* host name:       %s\n", HostName)
		fmt.Printf("************************************************************\n")

		//exit the app
		os.Exit(0)
	}
}
