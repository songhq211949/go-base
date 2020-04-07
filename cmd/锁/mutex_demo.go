package main

import (
	"fmt"
	"sync"
	"time"
)

//定义一个全局的互斥锁,初始化的时候它的状态为0
var synMutex sync.Mutex
//互斥锁 ,锁唯一且，不共享
func main() {
	go person1("hello")
	go person2("world")
	//不让主go程退出
	for {
		;
	}

}
//这里模拟两个同时使用打印机
func person1(str string){
	printer(str)
}
func person2(str string){
	printer(str)
}


func printer(str string){
	//这里上锁，表示正在使用打印机，别人就不能使用打印机了
	synMutex.Lock()
	//打印一个是逐字打印
	for _,v:= range str{
		fmt.Printf("%c",v)
		time.Sleep(time.Microsecond*300)//模拟打印的io是一个耗时的操作
	}
	synMutex.Unlock()
}
