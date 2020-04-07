package main

/**
函数是一等（first-class）类型
 */
import (
	"fmt"
	"strconv"
	"sync/atomic"
)
//员工ID生成器
type EmployeeIdGenerator func(company string,department string,sn uint32) string
//默认公司名称
var company = "baidu"
//序列号
var sn uint32
//生成员工ID
func generateId(generator EmployeeIdGenerator,department string)(string, bool){
	if generator == nil{
		return "",false
	}
	newSn := atomic.AddUint32(&sn,1)
	return generator(company,department,newSn), true
}
func appendSn(firstPart string,sn uint32)  string{
	return firstPart + strconv.FormatUint(uint64(sn),10)
}
func main() {
	var generator EmployeeIdGenerator
	generator = func(company string, department string, sn uint32) string {
		return appendSn(company+"-"+department+ "-",sn)
	}
	fmt.Println(generateId(generator,"RD"))
}




























