package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main(){
	//开启一个协程
	wg.Add(1)
	go func(){
		for i := 0; i < 10; i++{
			fmt.Println("hello go", i)
			time.Sleep(time.Millisecond * 100)
		}
		wg.Done()
	}()
	for i := 0; i <10; i++{
		fmt.Println("hello world", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
}