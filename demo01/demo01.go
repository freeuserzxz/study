package main

import "fmt"

func main() {
	// fmt.Println("hello golang")
	//fmt.Print("hello golang")
	// fmt.Printf("hello golang")
	/*
	Print和Println区别：
	1. Println会在输出内容后自动换行
	2. Println输出内容时，会对内容进行格式化
	3. Println输出内容时，会对内容进行转义
	*/
	fmt.Print("A","B","C")
	fmt.Println("A","B","C")//多句话有空格

	/*
	Println和Printf的区别：

	*/
	// //go语言变量定义了就必须使用
	// var a = "aaa"
	// fmt.Println(a)
	// //定义多个变量
	// var b,c = "bbb","ccc"
	// fmt.Println(b,c)

	// fmt.Printf("%v",a)

	var a int=10
	var b int=20
	fmt.Println(a,b)
	fmt.Println("a =",a,"b =",b)
	fmt.Printf("a = %v,b = %v\n",a,b)
	fmt.Printf("a = %v,b = %v\n",b,a)//一一对应 

	c:=30//类型推导方式定义变量
	fmt.Println(c)

	fmt.Printf("a=%v a的类型是%T",a,a)
}