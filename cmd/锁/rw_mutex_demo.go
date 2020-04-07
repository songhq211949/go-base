package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁   读锁共享  写锁不共享，且写锁比读锁优先级高
//而且 读写锁和channel最好不要混用

//定义全局变量
var value int =0
//定义读写锁
var rmMutex sync.RWMutex
func main() {
	//起5个读协程 起5个写协程
	for i := 0; i < 5; i++ {
		go readValue(i)
	}
	for i := 0; i < 5; i++ {
		go writeValue(i)
	}
	for true {
		;
	}

}
func writeValue(i int){
	for true {
		rmMutex.Lock()
		value++
		fmt.Printf("第读%dgo线程写为%d\n",i,value)
		time.Sleep(time.Microsecond*300)
		rmMutex.Unlock()

	}
}
func readValue(i int){
	for true{
		rmMutex.RLock()
		num := value
		fmt.Printf("------第读%dgo线程读到%d\n",i,num)
		time.Sleep(time.Microsecond*300)
		rmMutex.RUnlock()

	}
}
