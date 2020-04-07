package main

import (
	"bufio"
	"fmt"
	"go-project/pipeline"
	"net"
	"strings"
)

//简单的server服务器
func main01() {
	listen, err := net.Listen("tcp", ":7000")
	in := make(chan int,8)
	in <- 3
	in <- 4
	if err != nil{
		panic(err)
	}
	for{
		accept, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		writer := bufio.NewWriter(accept)
		pipeline.WriteSink(writer,in)

	}
}
//创建并发的服务器
func main(){
	lister,err := net.Listen("tcp","192.168.1.40:8000")
	if err != nil{
		fmt.Println("net.Listen is err",err)
		return
	}
	defer lister.Close()
	//监听客户端请求
	for {
		fmt.Println("服务器启动成功")
		con, err := lister.Accept()
		if err != nil{
			fmt.Println("lister.Accept is err",err)
			return
		}
		go handler(con)
	}
}
func handler(con net.Conn){
	defer con.Close()
	addr :=con.RemoteAddr()
	fmt.Println(addr,"客户端连接成功")
	data :=make([]byte,4096 )
	//循环读取客户端的数据
	for{
		n, err := con.Read(data)
		if err!=nil{
			fmt.Println("con.read is err",err)
			return
		}
		//打印客户端传过来的数据
		fmt.Println("客户端传过来的数据是", string(data[:n]))
		//这里完成小写转大写
		con.Write([]byte(strings.ToUpper(string(data[:n]))))
	}
}