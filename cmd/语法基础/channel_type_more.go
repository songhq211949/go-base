package main

import (
	"fmt"
	 "time"
	 )

type Sender chan <- int
type Receiver  <- chan int

func main() {
	var myChannel = make(chan int,0)//表示不带缓冲的channel
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Receive!",<- receiver)
	}()
	//让main函数执行结束的时间延迟1秒，以上两个代码块都有时间会被执行
	time.Sleep(time.Second)
}
