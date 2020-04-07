package main

import "fmt"

//百钱白鸡 公鸡 5钱 母鸡3钱  小鸡 三只1钱  100钱要买100只鸡
type name func()

var age int = 123

func main() {
	//test1(1)
	//匿名函数
	//v := func ( a int) int {
	//	fmt.Println(a)
	//	return 1
	//}(2)
	//
	//println(v)
	//闭包
	bao := biBao()
	//闭包的作用就是函数在调用后，函数内的变量还有效
	for i := 0; i < 10; i++ {
		i2 := bao()
		fmt.Println(i2)
	}
	//bao()

}
func test3() {

}
func modifySlice2() {

}
func main01() {}

func test() {

}
func forDemo() {
	//注意for循环的命名空间
	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			chicken := 100 - cock - hen
			count++
			if (cock*5+hen*3+chicken/3 == 100) && chicken%3 == 0 {
				fmt.Printf("公鸡：%d,母鸡：%d,小鸡%d \n", cock, hen, chicken)
			}
		}
	}
	fmt.Println("总共计算的次数为", count)

}
func test1(a ...int) {

	//test2(a) ,这里直接传a会报错 加...即可
	test2(a...)

}
func test2(a ...int) {
	for i, v := range a {
		println("element is", i, v)
	}
}
//闭包函数，函数的返回值是一个函数
func biBao() func() int{
	var a int
	return func() int {
		a++
		return a
	}
}