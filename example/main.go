package main

import (
	"flag"
	"fmt"

	"github.com/lovestreet/goversion" //自动打印 -v -version 启动参数并退出
)

var loglevel = flag.String("loglevel", "false", "the level of log")

func main() {
	flag.Parse()

	//[[ 代码集成
	//show version and exist, if have -v or -version
	goversion.ShowAndExist()
	//]]

	//do your code
	fmt.Printf("loglevel = %v", *loglevel)
}
