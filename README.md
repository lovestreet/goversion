# goversion
为 golang 应用程序编译加入版本信息

# 原理
使用 go build -ldflags "", 可以根据自己需要进行调整。

# 使用方法
1. 将example中的 Makefile copy 到 main 函数 所在的目录
2. 修改 Makefile 中的 目标生成文件名称
3. 每次发布时，手动调整 Makefile 中的版本信息
4. 在 shell 中执行 `make` or `make build`

# 注意事项
见 example 示例中的Makefile