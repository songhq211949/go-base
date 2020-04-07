package main

import "fmt"

func main() {
	student := new(Student)
	do(student)
	
}
type Behavior interface {
	Run()
	Eat()
}

//这里就是多态
func do(b Behavior)  {

	b.Eat()
	b.Run()

}
type Animal struct {
	Colour string
}
//定义一个学生类
type Student struct {
	Animal
	Name string
	Age string
}

func (a * Animal)Eat()  {
	fmt.Println("我的颜色是",a.Colour)

}

func (s *Student)Run()  {
	fmt.Println("is running")
}
