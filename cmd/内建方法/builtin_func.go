package main

import (
	"fmt"
	"reflect"
)

func main22() {

	//make创建 chan slice map  chan两个线程之间通信的
	//返回类型为引用类型
	//makeSlice()
	//makeMap()
	//makeChan()
	//new 内键方法返回 指针类型
	//newMap()
	//内键方append copy delete
	//modifySlice()
	//内键函数 panic recover
	//panicAndRecover()
	//
}

func panicAndRecover() {
	//哪怕捕获到了异常后面的代码也不会运营后面
	defer func() {
		message := recover()
		switch message.(type) {
		case string:
			fmt.Println("我是string异常")
		case error:
			fmt.Println("我是error")
		default:
			fmt.Println("我是 UnKnow ")
		}

	}()
	panic("i am panic")
	// panic(errors.New("我是异常"))
	// fmt.Println("我是panic后面要执行的数据")
}
func modifySlice() {
	strings := make([]string, 2) //make的时候已经初始化了
	strings[1] = "aa"
	fmt.Println(strings)
	fmt.Println(len(strings))
	fmt.Println(cap(strings))
	strings2 := append(strings, "bb")
	strings2[1] = "我是修改后的值"
	fmt.Println(len(strings2))
	fmt.Println(cap(strings2))
	//第一个是目标 第二个是源
	copy(strings, strings2)

	fmt.Println(strings)

	myMap := make(map[string]string, 1)
	myMap["1"] = "tom"
	myMap["4"] = "jim"
	delete(myMap, "1")
	fmt.Println(myMap)

}

func newMap() {
	myMap := new(map[string]string)
	myMap2 := make(map[string]string, 1)
	//打印什么类型
	fmt.Println(reflect.TypeOf(myMap))
	fmt.Println(reflect.TypeOf(myMap2))
}

func makeSlice() {
	mySlice := make([]string, 3)
	mySlice[0] = "a"
	mySlice[1] = "b"
	mySlice[2] = "c"
	//
	strings := append(mySlice, "d")
	fmt.Println(mySlice)
	fmt.Println(strings)
}
func makeMap() {
	myMap := make(map[string]string, 1)
	myMap["1"] = "tom"
	myMap["4"] = "jim"
	fmt.Println(myMap)

}

//这里的3 为缓存的大小为3 3删掉的话就没有缓存
func makeChan() {
	ints := make(chan int, 3)
	close(ints)

}
