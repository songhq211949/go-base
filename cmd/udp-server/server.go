package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//创建服务器的地址结构
	ipAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err!=nil {
		fmt.Println("创建udp对象失败",err)
		return
	}
	//创建通信的socket
	udpCon, err := net.ListenUDP("udp", ipAddr)//创建的时候不会有阻塞
	if err != nil{
		fmt.Println("创建udp连接失败",err)
		return
	}
	defer udpCon.Close()
	//读取客户端的数据
	buf := make([]byte, 4096)
	num, addr, err := udpCon.ReadFromUDP(buf) //直接阻塞在读上
	if err!=nil{
		fmt.Println("读取数据失败",err)
		return
	}
	//处理数据
	fmt.Printf("服务器读到 %v 的数据:%s \n",addr,string(buf[0:num]))

	//回显数据给客户端、
	 _,err = udpCon.WriteToUDP([]byte(time.Now().String()), addr)
	if err!=nil{
		fmt.Println("写数据失败",err)
	}

}
