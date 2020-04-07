package main

import (
	"encoding/json"
	"fmt"
)

//json的序列化和反序列化  Tag go中叫标签 对应java中的注解
func main() {
	//toJson()
	unSerialize()

}
type Student struct {
	Name string `json:"name"` //这里就是tag
	Age int `json:"age"`
}

func toJson(){
	s := new(Student)
	s.Age = 12
	s.Name = "tom"
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)//panic 为退出程序，如果没有recover()的话
	}
	fmt.Println("struct json序列化的结果是 ",string(b))

	m := make(map[string]interface{})
	m["age"] = 12
	m["name"] = "tom"
	c, err := json.Marshal(m)
	if err != nil {
		panic(err)//panic 为退出程序，如果没有recover()的话
	}
	fmt.Println("map json序列化的结果是 ",string(c))
}
func unSerialize(){
	var s Student
	jsonStr := ` {"name":"tom","age":12}`
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil{
		fmt.Println("发生了错误",err)
	}
	fmt.Println(s)

	m := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(jsonStr), &m)
	if err2 != nil{
		fmt.Println("发生了错误",err2)
	}
	fmt.Println(m)
}
