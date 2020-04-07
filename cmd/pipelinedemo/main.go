package main

import (
	"bufio"
	"fmt"
	"go-project/pipeline"
	"os"
)

func main() {

	source := pipeline.NetWorkReadSource(":7000")
	for v := range source{
		fmt.Println(v)
	}
	//mergeDemo()
	//outPut()
	//testNotCloseChan()
}
func testNotCloseChan()  {

	out := make(chan int)
	out <- 3
	for v := range out{
		fmt.Println(v)
	}
}
func outPut(){

	fileName := "small.out"
	n := 64
	file, err := os.Create(fileName)
	if err != nil{
		//panic不知道该怎么办
		panic(err)
	}
	//defer 函数退出前都会执行
	defer  file.Close()
	randomInt := pipeline.RandomInt(n)
	//使用缓冲流来包装普通的流
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer,randomInt)
	//使用buffer需要flush
	writer.Flush()

	//这里再把文件读出来
	open, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	resource := pipeline.ReadResource(bufio.NewReader(open),416)
	count := 0
	for v := range resource{
		fmt.Println(v)
		count++
		if count > 50 {
			//break
		}
	}
}

func mergeDemo() {

	//在这里就可以发现go语言的核心概念 就是channel 和 goroutine
	p := pipeline.MergeN(pipeline.SortInMem(pipeline.ArraySource(1,4,3,2)),
		pipeline.SortInMem(pipeline.ArraySource(1,4,3,2)))
	//for 死循环的写法
	//for{
	//	//从p中取数据
	//	if num, ok :=  <- p; ok {
	//		fmt.Println(num)
	//	}else{
	//		break
	//	}
	//}
	//利用range也可以实现遍历，不过这时候要close
	for num := range p {
		fmt.Println(num)
	}
}
