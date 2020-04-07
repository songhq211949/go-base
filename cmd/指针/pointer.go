package main

import "fmt"

//关于指针的要点  指针基本使用 指针数组 指向指针的指针 值传递和指针传递


func main() {
	//testPoint()
	//pointArray()
	st :=Student{12,"李四",12}
	point := &st
	fmt.Printf("结构体指针的数据类型%T\n",point)
	fmt.Printf("结构体变量的类型是 %T \n",&st)
	fmt.Printf("结构体变量的地址是 %p \n",&st)
}

//结构体指针
type Student struct {
	id int
	name string
	age int
}


func testPoint(){
	var count int = 20
	var countPoint *int
	//countPoint++ go语言中不支持指针的运算，没有指向具体变量的指针为野指针，go语言不支持
	countPoint = &count
	fmt.Println("count变量的地址",&count)
	fmt.Println("count变量的值",count)
	fmt.Println("countPoint 指针指向地址的变量值是",countPoint)
}
func pointArray(){
	a , b := 1,2
	pointArray := []*int{&a,&b}
	fmt.Println("指针数组是",pointArray)
	//没有指定大小的就是切片slice,go语言中优化了len()，和arrayPoint[0]取值,这种写法只有在数组指针时候支持
	//指向切片地额指针  属于二级指针

	array := []int{4,5,3}
	arrayPoint :=&array
	//for i,v:= range arrayPoint{
	//	fmt.Println(i,v)
	//
	//}
	fmt.Println((*arrayPoint)[0])
    fmt.Println(*arrayPoint)
	//fmt.Println(arrayPoint[0])

}