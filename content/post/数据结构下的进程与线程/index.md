# 数据结构下的进程与线程

进程是资源分配的最小单位, 对应一个虚拟内存空间
线程是CPU调度的最小单位, 对应的是一组CPU寄存器,描述了栈的状态
CS:IP, SS:SP对程序的输出是信息完备的
程序本质一套函数调用链.利用内存中的数据结构描述函数的规则, 调用CPU真正执行, 
在线程切换时, 保存CPU的寄存器, 因为

## go的并发之道
go语言并发采用fork-join模型
1G内存能创建30多万个goroutine