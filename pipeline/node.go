package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

//可变参数
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		//关闭chan
		close(out)
	}()
	return out
}

var startTime time.Time

func InitTime()  {
	startTime = time.Now()
}

//参数只入不进（针对函数本身而言），返回值只出不进
func SortInMem(in <-chan int) <-chan int {
	out := make(chan int,1024)
	go func() {
		//初始化一个slice,Read into memory
		var a []int
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read into memory done:" , time.Now().Sub(startTime))
		//Sort
		sort.Ints(a)
		fmt.Println("sort done:" , time.Now().Sub(startTime))
		//Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int,1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			{
				if !ok1 || v1 > v2 {
					out <- v2
					v2, ok2 = <-in2
				} else {
					out <- v1
					v1, ok1 = <-in1
				}
			}
		}
		fmt.Println("merge done:" , time.Now().Sub(startTime))
		close(out)
	}()

	return out
}

//文件输入，
func ReadResource(reader io.Reader,chunkSize int ) <- chan int {
	out := make(chan int,1024)

	go func() {
		//64位的计算机 int是8位
		buffer := make([]byte , 8)
		bytesRead := 0
		for{
			n,err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		//关闭chan
		close(out)
	}()
	return  out
}
//将排序的结果输出到文件中去
func WriteSink(writer io.Writer, in <- chan int){
	for v := range in {

		buffer := make([]byte , 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		fmt.Println("正在写入数据:", binary.BigEndian.Uint64(buffer))
		writer.Write(buffer)
	}
}
//输出一个chan
func RandomInt (count int)   <- chan  int {
	out := make(chan int)
	go func() {
		for i := 0;i<count;i++{
			data := rand.Int()
			out <- data
		}
		close(out)
	}()

	return out
}

//合并n个channel
func MergeN(inputs ...<- chan int ) <- chan int {
	if len(inputs) == 1{
		return inputs[0]
	}
	m := len(inputs)/2
	//这里使用到了递归算法
	return Merge(MergeN(inputs[:m] ...),MergeN(inputs[m:] ...))
}

