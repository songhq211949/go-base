package main

import "fmt"

type Int int

func (a Int) add (b Int) Int{
	return a+b
}
func add(a int,b int) int{
	return  a+b
}
//方法可以与函数的名称相同，在go语言中没有方法的重载，只要函数名称相同，就会产生冲突的
//func add(b string){
//	fmt.Println(b)
//}

//go的源码是分包管理的
func main01() {
	var f func()
	f = func() {
		fmt.Println("我是方法")
	}
	f()

	var a Int = 2
	b := a.add(3)
	fmt.Println(b)

}

type sayer interface {
	SayHello()
}
type Student struct {
	Hum
	name string

}
type Hum struct {
	sex string
}
func (s *Student) SayHello(){
	fmt.Println("我是学生")
}
func SayHello(sayer2 sayer){
	sayer2.SayHello()
}
func saySex(hu Hum){
	fmt.Println(hu.sex)
}

func main() {
	var say sayer
	say = &Student{}
	//子集不能接受超级  但是接口可以接受所有的实现类
	//saySex(say)
	say.SayHello()

	SayHello(&Student{})

}
