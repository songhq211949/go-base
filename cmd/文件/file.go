package main

import (
	"fmt"
	"os"
)

func main(){
	//查实申请非常大的栈内存
	a := [100000000]int{1,2,3}
	fmt.Println(len(a))

}

func main02(){
	fmt.Println(nil)
	fmt.Printf("%q\n",nil)//%q是打印go语言格式的字符串
}
func main01() {

	//创建摸个文件
	fp,err:=os.Create("D:/a.txt")
	if err!= nil {
		fmt.Println("创建文件错误")
		return
	}
	fp.WriteString("你好")
	defer fp.Close()
	
}
