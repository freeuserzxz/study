package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func judge(num int) {
	for i := (num-1)*30 + 1; i <= num*30; i++ {
		if i > 1 {
			flag := true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Println(i)
			}
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go judge(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
}
