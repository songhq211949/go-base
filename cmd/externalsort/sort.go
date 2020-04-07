package main

//外部排序即支持大文件排序

import (
	"bufio"
	"fmt"
	"go-project/pipeline"
	"os"
	"strconv"
)

func main() {
	//用网络写入读出失败了，是可以的，但是会丢失数据 坑
	p := createNetWorkPipeLine("large.in",80000000,4)
	//p := createPipeLine("small.in",512,4)
	writeFile(p,"large.out")
	printFile("large.out")

}

func printFile(fileName string) {
	open, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	p := pipeline.ReadResource(open, -1)
	count := 0
	for v := range p{
		fmt.Println(v)
		count++
		if count >=100 {
			break
		}
	}
}

func writeFile(p <- chan int,fileName string) {
	create, err := os.Create(fileName)
	if err != nil{
		panic(err)
	}
	defer create.Close()
	writer := bufio.NewWriter(create)
	defer writer.Flush()
	pipeline.WriteSink(writer,p)

}


//将一个文件分成几个部分然后合并成成一个chan
func createPipeLine(fileName string,fileSize, chunkCount int) <- chan int  {
	chunkSize := fileSize/chunkCount
	pipeline.InitTime()
	//定义一个chan 级别的slice,用来收集结果
	var sortResult []<-chan int
	for i := 0 ;i< chunkCount; i++ {
		open, err := os.Open(fileName)
		if err != nil{
			panic(err)
		}
		//分割文件
		resource := pipeline.ReadResource(bufio.NewReader(open), chunkSize)
		sortResult = append(sortResult,pipeline.SortInMem(resource))
		//fmt.Println(sortResult)

	}
	// ... 三个点表示什么
	return pipeline.MergeN(sortResult...)
}

//从网络中读写
func createNetWorkPipeLine(fileName string,fileSize, chunkCount int) <- chan int  {
	chunkSize := fileSize/chunkCount
	pipeline.InitTime()
	//定义一个chan 级别的slice,用来收集结果
	var sortAddr []string
	for i := 0 ;i< chunkCount; i++ {
		open, err := os.Open(fileName)
		if err != nil{
			panic(err)
		}
		//分割文件
		resource := pipeline.ReadResource(bufio.NewReader(open), chunkSize)
		addr :=  ":" + strconv.Itoa(i+7000)
		//将内存中排好序的写到server中
		pipeline.NetWorkSink(addr,pipeline.SortInMem(resource))
		//把写过的网络放在sortAddr中存储
		sortAddr = append(sortAddr,addr)
		//fmt.Println(sortResult)
	}
	//return nil
	var sortResult []<- chan int
	for _, v := range sortAddr{
		//从server中读取到channel
		ch := pipeline.NetWorkReadSource(v)
		sortResult = append(sortResult,ch)
	}
	// ... 三个点表示什么
	//最后是合并
	return pipeline.MergeN(sortResult...)
}