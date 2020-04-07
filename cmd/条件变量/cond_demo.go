package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//读写锁   读锁共享  写锁不共享，且写锁比读锁优先级高
//而且 读写锁和channel最好不要混用

//定义条件变量
var cond sync.Cond //这个里面有锁
func main() {
	//定义全局变量 channel
	channel := make(chan int, 4)
	//指定条件变量的锁
	cond.L = new(sync.Mutex)
	//起5个读协程 起5个写协程
	for i := 0; i < 5; i++ {
		go readValue(channel, i)
	}
	for i := 0; i < 5; i++ {
		go writeValue(channel, i)
	}
	for true {
		;
	}

}
func writeValue(out chan<- int, i int) {
	for true {
		cond.L.Lock()
		//判断当前channel是否满
		for ;len(out) == cap(out);  {
			cond.Wait()//这个会阻塞当前线程，且会解锁，唤醒时会加锁
		}
		value := rand.Intn(100)
		out <- value //使用了条件变量，就是没有用到channel自带的阻塞功能
		fmt.Printf("第读%dgo线程写为%d\n", i, value)
		cond.L.Unlock()
		cond.Signal()//环境对称的在等待的线程
		time.Sleep(time.Microsecond * 300)

	}
}
func readValue(in <-chan int, i int) {
	for true {
		cond.L.Lock()
		//判断当前channel是否满
		for ;len(in) == 0;  {
			cond.Wait()//这个会阻塞当前线程，且会解锁，唤醒时会加锁
		}
		num := <-in
		fmt.Printf("------第读%dgo线程读到%d\n", i, num)
		cond.L.Unlock()
		cond.Signal()//环境对称的在等待的线程
		time.Sleep(time.Microsecond * 300)
	}
}
