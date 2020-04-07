package main

import "fmt"

func main(){
	//声明一个整数类型变量并赋值
	var num int = -0x1000
	fmt.Printf("16进制数 %X 表示的是%d。\n",num,-4096)
}
