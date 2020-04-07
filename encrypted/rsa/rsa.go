package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	generateRsaKey(1024)
}
func generateRsaKey(size int){
	//rsa 会生成E N ,D N  ，私钥是在公钥里面
	//1.rsa的生成key rsa是三个开发者的姓氏
	key, err := rsa.GenerateKey(rand.Reader, size)
	if err!=nil{
		fmt.Println("加密失败",err)
	}
	//2.通过x509标准将得到的rsa秘钥序列化为pkcs1的DER编码的字符串  序列化为字节切片
	derText := x509.MarshalPKCS1PrivateKey(key)
	//3.构造一个Block对象
	block := pem.Block{
		Type:"rsa private key",
		Bytes:derText,
	}
	//4.pem编码
	//4.1创建文件指针
	file, err := os.Create("private.pem")
	if err!=nil{
		panic(err)
	}
	//4.2将秘钥写入文件中，是base64格式的
	err = pem.Encode(file, &block)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	// +++++公钥++++++
	//1、从私钥中取出公钥
	publicKey := key.PublicKey
	//2.使用x509标准序列化 PKIX，注意这里传入指针
	pkixPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}
	//3.构造一个Block对象
	publicBlock := pem.Block{
		Type:"rsa public key",
		Bytes:pkixPublicKey,
	}
	//4.pem编码
	//4.1创建文件指针
	publicFile, err := os.Create("public.pem")
	if err!=nil{
		panic(err)
	}
	//4.2将公钥写入文件中，是base64格式的
	err = pem.Encode(publicFile, &publicBlock)
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()

}
