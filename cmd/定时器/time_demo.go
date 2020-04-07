package main

import (
	"fmt"
	"time"
)

func main() {
	//有这样一个Timer类
	//type Timer struct {
	//	C <-chan Time
	//	r runtimeTimer
	//}
	fmt.Println("系统当前时间为",time.Now())
	newTimer :=time.NewTimer(time.Second*2)//系统会在两秒后向Timer中的C通道写入当前的时间
	nowTime := <-newTimer.C //读2秒后后的时间
	fmt.Println("现在的时间为",nowTime)

	
}
