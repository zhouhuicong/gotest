package main

import (
	// "errors"

	"context"
	"fmt"
	"sync"
	"time"
)

type Persion struct {
	name string
	age  int
}
type Xingwei interface {
	Walk()
	Run()
}

func (p Persion) Walk() {
	fmt.Println(p.name + " is walk")
}

func (p Persion) Run() {
	fmt.Println(p.name + " is run")
}

func downLoadFile(filename string) string {
	time.Sleep(time.Second * 1)
	return "filePaht:" + filename
}

//共享资源
var sum int
var mutex sync.RWMutex

//
func main() {
	// makeSlice := []int{1, 3, 20, 4, 7, 2}
	// sort.Ints(makeSlice)
	// fmt.Println(makeSlice)
	// fmt.Println(makeSlice)
	// fmt.Println(makeSlice)
	// e := errors.New("this is a new error")
	// defer fmt.Println("the firest")
	// defer fmt.Println("the second")
	// defer fmt.Println("the third ")
	// e := errors.New("原始错误e")
	// w := fmt.Errorf("Wrap了一个错误:%w", e)
	// fmt.Println(w)
	// fmt.Println(errors.Unwrap(w))
	// fmt.Println(errors.Is(w, e))
	// ch := make(chan string)
	// go func() {
	// 	fmt.Println("this is a goroutine")
	// 	ch <- "this goroutine is complete"
	// }()

	// fmt.Println("i am the main goroutine")
	// v := <-ch
	// fmt.Println(v)
	// fmt.Printf("this  other goroutine message: %T %v\n ", errors.Cause(e), errors.Cause(e))
	// fmt.Println("---------------------")
	// fmt.Printf("this  other goroutine message:%+v", errors.Cause(e))
	// firstCh := make(chan string)
	// sencondCh := make(chan string)
	// thirdCh := make(chan string)
	// go func() {
	// 	firstCh <- downLoadFile("theFirst")
	// }()
	// go func() {
	// 	sencondCh <- downLoadFile("theSecond")
	// }()
	// go func() {
	// 	thirdCh <- downLoadFile("theThird")
	// }()

	// select {
	// case first := <-firstCh:
	// 	fmt.Println(first + "  is complete")
	// case second := <-sencondCh:
	// 	fmt.Println(second + " is complete")
	// case third := <-thirdCh:
	// 	fmt.Println(third + " is complete")
	// 	// default:
	// }
	// fmt.Println(time.Now().Clock())
	// test()
	// fmt.Println(time.Now().Clock())
	// a := func() {
	// 	fmt.Println("this is niminghanshu")
	// }
	// a()
	// nimingtest(func(g string) { fmt.Println("print ", g) })
	// race()
	// testMap("1", "1")
	// testMap("2", "22")
	// testMap("3", "333")
	var wg sync.WaitGroup
	wg.Add(3)
	// stopCh := make(chan bool)
	// go func() {
	// 	watchDog(stopCh, "zhc")
	// 	wg.Done()
	// }()
	// time.Sleep(time.Second * 5)
	// //停止监控
	// stopCh <- true

	ctx, stop := context.WithCancel(context.Background())
	ctxz, stopz := context.WithCancel(ctx)
	go func() {
		defer wg.Done()
		watchDog1(ctx, "zhc1")
	}()
	go func() {
		defer wg.Done()
		watchDog1(ctxz, "zhc2")
	}()
	go func() {
		defer wg.Done()
		ctxtemp := context.WithValue(ctxz, "username", "zhctest")
		getUsername(ctxtemp)
	}()
	time.Sleep(time.Second * 5)
	stopz() //子类发出停止指令
	time.Sleep(time.Second * 3)
	stop() //父类停止
	wg.Wait()
}
func getUsername(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("this goroutine is complete")
		return
	default:
		fmt.Println("username is ", ctx.Value("username"))
		time.Sleep(time.Second * 1)
	}
}
func watchDog(stopCh chan bool, name string) {
	//开启for select模式
	for {
		select {
		case <-stopCh:
			fmt.Println("the watching is stop")
			return
		default:
			fmt.Println(name, " is watching ")
		}
		time.Sleep(time.Second * 1)
	}
}
func watchDog1(ctx context.Context, name string) {
	//开启for select模式
	for {
		select {
		case <-ctx.Done():
			fmt.Println("the watching is stop")
			return
		default:
			fmt.Println(name, " is watching ")
		}
		time.Sleep(time.Second * 1)
	}
}

// var testmap sync.Map

// func testMap(k, v string) {

// 	testmap.Store(k, v)
// 	testmap.Range(func(key, value interface{}) bool {
// 		fmt.Println(key, "  ", value)
// 		return true
// 	})
// 	// fmt.Println(testmap)
// }

// func nimingtest(f func(string)) { f("zzzz") }
// func SumCount(i int) {
// 	mutex.Lock()
// 	sum = i + sum
// 	mutex.Unlock()
// }
// func ReadSum() int {
// 	mutex.RLock()
// 	r := sum
// 	mutex.RUnlock()
// 	return r
// }
// func test() {
// 	var wg sync.WaitGroup
// 	wg.Add(1100)

// 	for i := 1; i < 1001; i++ {
// 		go func() {
// 			defer wg.Done()
// 			SumCount(10)
// 		}()

// 	}

// 	for i := 1; i < 101; i++ {
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("当前的sum值为: ", sum)
// 		}()

// 	}
// 	wg.Wait()

// }

// func race() {
// 	cond := sync.NewCond(&sync.Mutex{})
// 	var wg sync.WaitGroup
// 	wg.Add(11)
// 	for i := 1; i < 11; i++ {
// 		go func(num int) {
// 			//runer就位
// 			fmt.Println(num, " is ready")
// 			cond.L.Lock()
// 			cond.Wait()
// 			fmt.Println(num, " is running")
// 			cond.L.Unlock()
// 			wg.Done()
// 		}(i)
// 	}
// 	//等待所有的goroutine 进入wait
// 	time.Sleep(time.Second * 2)
// 	go func() {
// 		fmt.Println("裁判已经就位")
// 		fmt.Println("比赛开始")
// 		cond.Broadcast() //发令枪响
// 		wg.Done()
// 	}()
// 	//等待计数器为0
// 	wg.Wait()
// }
