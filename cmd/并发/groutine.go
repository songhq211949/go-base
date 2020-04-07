package main

import (
	"fmt"
	"sync"
	"time"
)
var wait sync.WaitGroup


func main(){
	//使用channel

	channel := make(chan int)
	//没有用到channel的读阻的时候
	go func() {
		//channel <- 123
		person1("hello") //这样直接使用channel打印的东西会串
		channel <- 123//写的话也会阻塞，知道有人读的时候才会取消阻塞 channel的使用降低了锁的难度
	}()
	go func() {
		<- channel //这里的读会阻塞
		person2("world") //这样直接使用channel打印的东西会串
	}()
	for {
		;
	}
}
//go语言并发演示
func main01() {
	//go 关键字启动类似线程的协程
	//go printInt()
	//go printInt()
	//
	//time.Sleep(time.Second)

	//线程的同步
	read()
	go Write()
	wait.Wait()
	fmt.Println("我是主程序，理应在写完毕后才打印")
	time.Sleep(time.Second)
}

func printInt(){
	for i := 0; i<5 ; i++  {
		//microsecond 微妙
		time.Sleep(time.Microsecond * 10)
		fmt.Println("输出了：",i)
	}
}
//线程中通信使用的是channel
//线程中的同步使用的 sync.WaitGroup 这个工具
func read()  {
	for i := 0 ; i< 6 ; i++ {
		wait.Add(1)
	}
}
func Write()  {
	for i := 0 ; i< 6 ; i++ {
		time.Sleep(time.Microsecond * 10)
		wait.Done()
	}
	fmt.Println("写完毕了")
}

//可利用channel的无缓冲的读阻塞来实现先后顺序
func person1(str string){
	for _,v := range str{
		time.Sleep(time.Microsecond*300)
		fmt.Printf("%c",v)
	}
}
func person2(str string){
	for _,v := range str{
		time.Sleep(time.Microsecond*300)
		fmt.Printf("%c",v)
	}
}




