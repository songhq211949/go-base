package main

import "fmt"

/**
一个数组（Array）就是一个可以容纳若干类型相同的元素的容器。这个容器的大小（即数组的长度）是固定的，且是体现在数组的类型字面量之中的
 */
func main(){
	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
}
