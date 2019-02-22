#******************************************************************************
# golang application build version with Makefile
#
# 文件位置：与main.go同目录
# 执行方法：在shell中执行 make build
# 检查方法：在shell中执行 ./{application} -v 或者 ./{application} -version
#
# BUILD_NAME :      生成的APP名称,一般与main.go同目录,需要手动设置
# BUILD_VERSION :   手动配置的程序版本，需要手动设置
# BUILD_TIME :      执行编译的时间，自动收集
# GIT_COMMIT :      执行编译时的 git 中的最后一条commit，自动收集
# GIT_BRANCH :      执行编译时的 git 分支，自动收集
# GOLANG_VERSION :  golang 的版本信息，自动收集
# HOST_NAME :       当前编译主机的名称，自动收集
# 
#******************************************************************************
BUILD_NAME := showVersion
BUILD_VERSION := 1.1.1 
BUILD_TIME := `date "+%F %T"`
GIT_COMMIT := `git rev-parse HEAD` 
GIT_BRANCH := `git branch | grep \* | cut -d ' ' -f2` 
GOLANG_VERSION := `go version` 
HOST_NAME := `hostname` 

build: 
	# build the app
	go build -ldflags                                                   \
	"                                                                   \
	-X 'github.com/lovestreet/goversion.BuildVersion=${BUILD_VERSION}'  \
	-X 'github.com/lovestreet/goversion.BuildTime=${BUILD_TIME}'        \
	-X 'github.com/lovestreet/goversion.BuildName=${BUILD_NAME}'        \
	-X 'github.com/lovestreet/goversion.GitCommitID=${GIT_COMMIT}'      \
	-X 'github.com/lovestreet/goversion.GitBranch=${GIT_BRANCH}'        \
	-X 'github.com/lovestreet/goversion.GolangVersion=${GOLANG_VERSION}'\
	-X 'github.com/lovestreet/goversion.HostName=${HOST_NAME}'          \
	"                                                                   \
	 -o ${BUILD_NAME}

	# output the version
	./${BUILD_NAME} -v

clean: 
	rm -f ${BUILD_NAME} 


.PHONY : build clean ${BUILD_NAME}
