package main

import (
	"fmt"
	"sync"
	"time"
)

func buy(num int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i < num+1; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}
func zuzhuang(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for c := range in {
			out <- "组装[" + c + "] "
			time.Sleep(time.Millisecond * 1)
		}
	}()
	return out
}

func baozhuang(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "包装【" + c + "】"
		}
	}()
	return out
}
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(ins))
	out := make(chan string)

	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	for _, in := range ins {
		go p(in)
	}
	go func() {
		wg.Wait()
		fmt.Println("out  is close")
		close(out)
	}()

	return out
}
func main() {
	c := buy(100)
	zh1 := zuzhuang(c)
	zh2 := zuzhuang(c)
	zh3 := zuzhuang(c)
	// for z := range zh3 {
	// 	fmt.Println(z)
	// }
	zh := merge(zh1, zh2, zh3)
	bz := baozhuang(zh)
	for b := range bz {
		fmt.Println(b)
	}
}
