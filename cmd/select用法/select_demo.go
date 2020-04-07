package main

import "fmt"

func printData(data <-chan int,flag <-chan bool){

	//这种是阻塞轮训 ,如果加上了default则是忙轮训，select需结合for循环使用
	for  {
		select {
			case num := <- data:
				fmt.Print(num," ")
			case <-flag:
				return
		}

	}
}
//实现斐波那契数列  1 1 2 3 5 ...
func main() {
	data := make(chan int)
	flag := make(chan bool)
	x,y := 1,1
	go printData(data,flag)//打印数列
	for i := 0; i < 20; i++ {
		data <-x //这里没有用到缓冲，会写一个读一个
		x,y = y,x+y
	}
	flag  <- true
}
