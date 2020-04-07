package main

import "fmt"

func main01() {
	//使用中文时会overflow

	var a byte = 'a'
	var b string = `hello\0`
	c := "我心飞翔"  //一个汉字算作是三个字符
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%s",b)
	fmt.Println("b字符串的长度为",len(b))
	fmt.Println("b字符串的长度为",len(c))

}
func main() {
	var name string
	fmt.Print("请输入你的分数")
	var score int
	//fmt.Scanf("%d",&score)
	if score>900 {
		fmt.Println("我要上清华")
	}else if score>500{
		fmt.Println("我要上二本")
	}else if name == ""{
		fmt.Print("要么牛逼，要么菜的坑逼")
	}
	fmt.Println(name == "")
}