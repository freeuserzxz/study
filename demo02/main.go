package main
import "fmt"
func main(){
	//声明变量，字母、数字、下划线组成，但不能以数字开头，不能使用关键字和保留字
	var username string
	username = "zhangsan"
	fmt.Println(username)

	var name = "tom"
	fmt.Println(name)
	var age = 10
	var sex = "男"
	fmt.Println(name,age,sex)
	
	var a1,a2 string
	a1 = "aaa"
	a2 = "aaaaa"
	fmt.Println(a1,a2)

	var(
		a3 string
		a4 int
		a5 bool
	)
	a3 = "aaa"
	a4 = 10
	a5 = true//也可以定义的时候一起赋值，不写类型，会自动推导
	fmt.Println(a3,a4,a5)

	//短变量声明，只能在函数内部使用
	name1 := "tom"
	age1 := 10
	fmt.Println(name1,age1)
	//匿名变量，用_表示，不占用内存空间，也不会报错
	//比如函数返回两个值，我们只需要一个值，另一个值用_接受即可

}