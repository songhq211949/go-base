package main

import (
	"fmt"
	"net"
)
//创建用户结构体
type User struct {
	C chan string
	Name string
	Addr string
}
//创建全局map的map,储存上线用户，此时还未初始化
var onlineMap map[string]User
//创建全局channel
var channelMsg =make(chan string)  //这个全局map是否为线程安全的,答案是否定的


//处理用户连接
func handlerConnect(con net.Conn)  {
	//defer con.Close()
	//创建一个用户体,每个用户创建一个channel
	remoteAddr := con.RemoteAddr().String()
	user := User{make(chan string),remoteAddr,remoteAddr}
	fmt.Println("用户进来了",user)
	//向map中添加用户
	onlineMap[remoteAddr] = user
	//创建专门为该用户创建的线程
	go writeMsg(con,user.C)
	//发送用户上线消息到广播中
	channelMsg <-"[" + user.Name +"]"+user.Addr +"login"
	//for{
	//	;
	//}
}

//写消息给用户
func writeMsg(con net.Conn, c chan string) {
	for msg := range c{
		//这个“\n”是刷新缓冲的作用
		con.Write([]byte(msg + "\n"))
	}
}
//创建管理者
func manager(){
	//初始化map
	onlineMap = make(map[string]User)
	for{
		//监听全局channel是否有msg
		msg := <- channelMsg
		//然后去循环发送消息给在线用户
		for _,value := range onlineMap{
			value.C <- msg
		}
	}

}


func main() {
	lister,err:=net.Listen("tcp",":8089")
	if err !=nil{
		fmt.Println("net.Listen",err)
	}
	defer lister.Close()

	//创建管理者，管理map和全局channel
	go manager()
	for{
		con,err :=lister.Accept()
		if err != nil {
			fmt.Println("lister.Accept",err)
		}
		handlerConnect(con)
	}

}

