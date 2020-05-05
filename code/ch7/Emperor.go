package main

import "fmt"

//定义皇帝结构体
type Emperor struct {
}

//定义说话方法
func (E *Emperor) say() {
	fmt.Println("我就是皇帝某某某")
}

//懒汉在初始化时进行实例
var emperor *Emperor

func init() {
	fmt.Println("初始化一个皇帝")
	emperor = new(Emperor)
}

//获取实例
func getInstance() *Emperor {
	return emperor
}

//实现多人对一皇帝叩拜
func main() {
	for i := 0; i < 10; i++ {
		e := getInstance()
		e.say()
	}
}
