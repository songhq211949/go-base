package main

import "fmt"
func main(){
	//声明一个string类型变量并赋值
	var str1 string = "\\\""
	//这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹,表现值，并不是解析后的值
	fmt.Printf("用解释型字符串表示法表示的%q所代表的是 %s",str1,str1)
}
