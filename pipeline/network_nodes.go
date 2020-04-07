package pipeline

import (
	"bufio"
	"fmt"
	"net"
)

//单机上 goroutine 与 goroutine 是通过channel来通信的
//网络上 goroutine --> WriterSink 网络-->ReadSource -->goroutine
// InMemorySort -->WriterSink 网络-->ReadSource -->goroutine

//往网络写,从一个channel的数据写到网络中去
func NetWorkSink(addr string,in <- chan int){
	listen, err := net.Listen("tcp", addr)
	if err != nil{
		 panic(err)
	}
	//开完server,程序就立马执行别的东西了
	go func() {
		defer  listen.Close()
		accept, err := listen.Accept()
		//阻塞在这里了
		if err != nil {
			panic(err)
		}
		defer accept.Close()
		writer := bufio.NewWriter(accept)
		WriteSink(writer,in)
	}()

}
//从网络中读数据,把数据读到 channel中去，这个方法根本就娶不到数据
func NetWorkReadSource(addr string)  <- chan int{
	fmt.Println("尝试从server中读取数据" + addr)
	p := make(chan  int)
	go func() {
		//连接到网络
		dial, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		//defer dial.Close()
		resource := ReadResource(bufio.NewReader(dial), -1)
		for v := range resource{
			fmt.Println("读到数据了 啊啊啊")
			p <- v
		}
		close(p)
	}()
	return  p
}

