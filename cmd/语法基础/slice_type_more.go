package main

import "fmt"

func main(){
	var num = [...]int{1,2,3,4,5,6}
	slice := num[2:3:6]
	fmt.Printf("slice的数据大小为为%d",len(slice))
	fmt.Printf("slice的容量为%d",cap(slice))
}
