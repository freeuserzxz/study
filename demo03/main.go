package main

import "fmt"

func main() {
	// var username = "zhangsan"
	// username = "lisi"
	// fmt.Println(username)

	//常量,定义的时候必须赋值，且不能修改
	const pi = 3.14159
	const (
		a = 1
		b = "B"
		c
		d
		//省略了值的时候和上一行是一样的
	)
	fmt.Println(pi)
	fmt.Println(a, b, c, d) //1 B B B

	//iota的使用,常量计数器，const中每新增一行常量声明iota计数一次
	//1.每次const出现的时候，iota都会重置为0
	const A = iota
	const B = iota
	fmt.Println(A, B) //0 0
	//2.每一行的iota都会增加1
	const (
		C = iota
		D
		E
	)
	fmt.Println(C, D, E) //0 1 2
	//3.如果中间插队，iota也会重新计算
	const (
		F = iota
		G = 100
		H
		I = iota
	)
	fmt.Println(F, G, H, I) //0 100 100 3
	//4.多个iota定义在一行
	const (
		J, K = iota, iota + 1
		L, M
		N, P
	)
	fmt.Println(J, K, L, M, N, P) //0 1 1 2 2 3

}
