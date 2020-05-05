package main

import (
	"fmt"
	"sync"
	"time"
)

//定义单例
type Singleton struct {
}
//定义说话方法
func (S *Singleton) say() {
	fmt.Println("我只创建了一次")
}

//懒汉在初始化时进行实例
/*
var singlenton  = new(Singleton)
//获取实例
func getInstance1() *Singleton {
	return singlenton
}
*/
/*
var singlenton  *Singleton
//获取实例
func getInstance2() *Singleton {
	if singlenton == nil {
		fmt.Println("饿汉初始化创建")
		singlenton = new(Singleton)
	}
	return singlenton
}
*/

//获取实例
var singlenton  *Singleton
var lock *sync.Mutex = &sync.Mutex{}

func getInstance3() *Singleton {
	if singlenton == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("饿汉初始化创建")
		singlenton = new(Singleton)
	}
	return singlenton
}
func main() {
	go fmt.Printf("%p\n", getInstance3())
	go fmt.Printf("%p\n", getInstance3())
	go fmt.Printf("%p\n", getInstance3())
	time.Sleep(time.Second)
	fmt.Println("done")
}
