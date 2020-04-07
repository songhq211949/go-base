package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "192.168.1.40:8000")
	if err!=nil{
		fmt.Println("net.dial is err",err)
	}
	defer con.Close()
	go write(con)
	for{
		data := make([]byte,4096)
		n,err :=con.Read(data)
		if err !=nil{
			fmt.Println("con.read is err",err)
			return
		}
		fmt.Println("读到服务器的数据为",string(data[:n]))
	}


}
func write(con net.Conn){
	data := make([]byte,4096)
	for true {
		//从键盘读取标准输入
		n, err2 := os.Stdin.Read(data)
		if err2 != nil{
			fmt.Println("标注输入错误")
		}

		str :=string(data[:n])
		if str == "exit\r\n" || str == "exit\n"{
			con.Close()
			return
		}
		_, err := con.Write(data[:n])
		if err !=nil{
			fmt.Println("con.read is err",err)
			return
		}
	}
}