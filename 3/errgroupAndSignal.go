package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sync/errgroup"
)

func GracefullExit(stop context.CancelFunc) {
	fmt.Println("Start Exit...")
	fmt.Println("stop  the http ...")

	fmt.Println("End Exit...")
	stop() //停止http server
	os.Exit(0)
}

func ServerApp(ctx context.Context, addr string, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// time.Sleep(5 * time.Second)
		fmt.Fprintln(resp, "this is a zhouhuicong http go test ")
	})
	server := &http.Server{
		Addr:    fmt.Sprint(addr, ":", port),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	select {
	case <-ctx.Done():
		fmt.Println("停止指令已收到，http 马上停止")
		server.Shutdown(ctx)
		return
	default:
		fmt.Println("程序运行中")
		time.Sleep(time.Second)

	}

}

//linux signal
func signalFun(stop context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Program Exit...", s)
				GracefullExit(stop)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()
}

//校验是否有协程已发生错误
func CheckGoroutineErr(errContext context.Context) error {
	select {
	case <-errContext.Done():
		fmt.Println("报错啦")
		return errContext.Err()
	default:
		fmt.Println("没错误")
		return nil
	}
}
func main() {

	fmt.Println("Program Start...")
	//创建 context
	ctx, stop := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)
	for index := 0; index < 3; index++ {
		indexTemp := index // 子协程中若直接访问index，则可能是同一个变量，所以要用临时变量

		// 新建子协程
		g.Go(func() error {
			// fmt.Printf("indexTemp=%d \n", indexTemp)
			if indexTemp == 0 {
				fmt.Println("http start ")
				ServerApp(ctx, "127.0.0.1", 8080)
			} else if indexTemp == 1 {
				fmt.Println("signal start")
				signalFun(stop)
			} else if indexTemp == 2 {
				fmt.Println("捕获子协程error begin")

				// 用于捕获子协程的出错, wait for stop siganl
				<-errCtx.Done()

				//检查 其他协程已经发生错误，如果已经发生异常，则不再执行下面的代码
				err := CheckGoroutineErr(errCtx)

				if err != nil {
					fmt.Println("子协程error,准备退出")
					GracefullExit(stop) // os 和 http stop
					return err
				}
			}
			return nil
		})
	}
	// 捕获err
	err := g.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Printf("get error:%v", err)
	}

}
