package main

import (
	"fmt"
)

// 主要为 ... 和 :=
func main() {
	sugar("你好","我也好")
	
}
func sugar(values... string)  {
	for _,v := range values{
		fmt.Println("值为",v)
	}
}
