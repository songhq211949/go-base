package main

import "fmt"
/**
  切片（Slice）与数组一样，也是可以容纳若干类型相同的元素的容器。与数组不同的是，无法通过切片类型来确定其值的长度。每个切片值都会将数组作为其底层数据结构。我们也把这样的数组称为切片的底层数组
 */
func main() {
	//var array1 = [5]int{1,2,3,4,5}
	//slice1 := array1[2:len(array1)]
	//slice2 := slice1[1:cap(slice1)]
	//fmt.Printf("%v ,%v,%v \n",3 == len(slice1),3 == cap(slice1),3 == len(slice2))
	//原始的值会改变 ，底层因为切片地存储地址是在堆，与数组的存储地址不同

	s := []int{1,2,3,4}
	changeSlice(s)
	fmt.Println(s)

}
func changeSlice(a []int){
	a[3] = 100
}