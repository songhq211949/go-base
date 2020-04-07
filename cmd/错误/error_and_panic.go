package main

import (
	"errors"
	"fmt"
)

//这个方法返回error
func err() (value int,error error){
	value=1
	error=errors.New("我发生的错误")
	return
}
//这个方法会发生异常
func happenPanic(){
	defer func() {
		i := recover()
		fmt.Println(i)
		a:=1
		b:=0
		c:=a/b
		println(c)
	}()
	b:=1
	c:=1
	a := b/c
	fmt.Println(a)
}

func main() {
	fmt.Println("1")
	err()
	fmt.Println("2")
	happenPanic()
	fmt.Println("3")
}
