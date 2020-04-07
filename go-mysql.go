package src

import (
	"database/sql"//这个是go sdk中
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_下换下线是为了不报错，这个只是驱动，只有运行的时候不报错
)
func main(){
	//构建go项目的时候时候可以采用传统的GOPATH的方式，也可以采用go mod模块的方式  go mod是go 1.11版本才有的
	//连接数据库,注意括号
	db, err := sql.Open("mysql", "root:root@(songhq.club:3306)/him")
	defer db.Close()
	if err != nil{
		fmt.Println("连接数据库错误",err)
		return
	}
	err= db.Ping()//这个才是正在的连接数据库
	if err !=nil{
		fmt.Println("数据库连接失败",err)
	}


}


