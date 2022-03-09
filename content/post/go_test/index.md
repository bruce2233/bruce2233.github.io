+++

author = "YunPeng Zhang"
title = "单元测试神器之go test"
date = "2022-03-08"
categories = [
"Go",

"测试"

]
tags = [
"Test"
]

image = "goimg.jpg"

+++

​                                                            [![go-test](https://img.shields.io/badge/testing-api-brightgreen?style=flat&logo=go)](https://studygolang.com/pkgdoc)                [![go-test](https://img.shields.io/badge/go-testing-brightgreen?style=flat&logo=github)](https://github.com/golang/go/tree/master/src/testing)

`go test` 是 Go 标准库中的`单元测试`与性能测试系统,也是 go 的`子命令`之一

## go test 命令及常用参数

`go test [文件名] [-run 函数名]`执行工作目录下`[文件名]`内`函数名`

- 测试文件名**必须** 以`_test.go`结尾
- 测试函数名**必须**以`Test`开头`&&` 参数类型必须为 `*tesing.T`
- 若测试代码引用包内其他`.go`文件，**必须** 在命令中添加
- `[文件名] == nil` ==工作目录下所有文件
- `-run 函数名 == nil` == 工作目录下所有文件

`-v` 输出详细测试信息 

`-bench`基准测试参数，待更新......
