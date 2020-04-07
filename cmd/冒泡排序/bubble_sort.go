package main

import "fmt"

//冒泡排序
func main() {
	arr := [10]int{1,25,7,5,6,4,3,4,3,4}
	fmt.Printf("arr的数据类型为%T\n",arr)
	//go语言中是值传递
	bubbleSort(arr)
	fmt.Println(arr)
	
}
//冒泡排序 相邻的数之间比较，然后交换 ，第一趟全部遍历，第二趟遍历到倒数第二，依次类推
func bubbleSort(arr [10]int){
	for i:=0;i<len(arr)-1 ;i++  {
		for j:= 0; j<len(arr)-1-i;j++  {
			if arr[j]> arr[j+1] {
				//go语言特有的交换值的写法
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}
