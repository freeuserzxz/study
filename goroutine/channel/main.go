package  main
import (
	"fmt"
)
func main() {
	ch := make(chan int,3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)
	/*
	for range会一直阻塞等待从通道接收数据
	关闭通道后，for range能感知到通道关闭，自动退出循环
	*/
	for val := range ch{// channel没有key，只遍历value
		fmt.Println(val)
	}
	/*
	chan<- type 为只写通道
	<-chan type 为只读通道
	用在函数里面，作为参数
	*/
}