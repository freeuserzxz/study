package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	//1.定义int类型
	var num int8 = 10
	fmt.Printf("num = %v 类型:%T\n",num,num)
	fmt.Println(unsafe.Sizeof(num))

	//int类型转换,高转低注意溢出问题，int8是-128~127
	var num1 int8 = 10
	var num2 int32 = 10
	fmt.Println((int32(num1)+num2))

	//2.浮点型
	var num3 float32 = 1.1
	fmt.Printf("num3 = %v 类型:%T\n",num3,num3)
	fmt.Printf("num3 = %f 类型:%T\n",num3,num3)
	fmt.Println(unsafe.Sizeof(num3))

	var c float64 = 3.1415926535
	fmt.Printf("c = %v\n",c)
	fmt.Printf("c = %f\n",c)
	fmt.Printf("c = %.2f\n",c)
	fmt.Printf("c = %.4f\n",c)

	//3.科学计数法
	var num4 float32 = 1.1e2  //表示1.1*10^2
	fmt.Printf("num4 = %v 类型:%T\n",num4,num4)
	fmt.Printf("num4 = %f 类型:%T\n",num4,num4)
	num4 = 1.1e-2  //表示1.1*10^-2
	fmt.Printf("num4 = %v 类型:%T\n",num4,num4)
	fmt.Printf("num4 = %f 类型:%T\n",num4,num4)

	var c1 float64 = 23.65
	c2 := int(c1)
	fmt.Printf("c2 = %v 类型:%T\n",c2,c2)

	str:=`hello
	world
	定义多行字符串
	用反引号`
	fmt.Println(str)

	var str1 string ="123+456+789"
	str2 := strings.Split(str1,"+")
	fmt.Println(str2)
	
	str3 := strings.Join(str2,"+")
	fmt.Println(str3)

	str4 := "this is a test"
	str5 := "test"
	flag := strings.Contains(str4, str5)
	fmt.Println(flag)
}