package main

import "fmt"

func main() {
	//TestForStruct()
	testMultiExtend()
}
func testMultiExtend(){

	var c C//声明变量是不会包空指针的 但如果指针变量初始值可能会导致nil空指针

	//这里会引起歧义 ambiguous

	c.A.name ="张三"

	fmt.Println(c)

}




type A struct {
	name string
	age int
}
type B struct {
	name string
	sex string
}
type C struct {
	A
	B
	score int
}

type Animal struct {
	Colour string
}
//定义一个学生类
type Student struct {
	Animal
	Colour int
	Name string
	Age string
}

func (a * Animal)Eat()  {
	fmt.Println("我的颜色是",a.Colour)

}

func (s *Student)Run()  {
	fmt.Println("is running")
}


func TestForStruct()  {
	//第一种 var来声明
	var student Student
	student.Age = "26"
	student.Name = "我是tom"
	fmt.Println("student:",student)

	//第二种
	student1 := Student{Name:"A",Age:"12"}
	fmt.Println(student1)

	//第三种,这种返回的是指针类型
	student2 := new(Student)
	student2.Name = "B"
	student2.Age = "13"
	student2.Animal.Colour = "red"
	student2.Run()
	student2.Eat()
	fmt.Println(student2)
	
}

