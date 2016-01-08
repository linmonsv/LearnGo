package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func exmaple_goroutine() {
	//默认情况下，调度器仅使用单线程，也就是说只实现了并发
	go say("world")
	say("hello")

	//想要发挥多核处理器的并行，需要在我们的程序中显示的调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。
	//GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。
	//如果n < 1，不会改变当前设置。
	//以后Go的新版本中调度得到改进后，这将被移除
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

/*
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，
这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
无缓冲channel是在多个goroutine之间同步很棒的工具
*/
func exmaple_channel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func example_Buffered_Channels() {
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func exmaple_range_close() {
	fmt.Println("exmaple_range_close")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			//当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）
		}
	}
}

func exmaple_select() {
	fmt.Println("exmaple_select")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func example_timeout() {
	fmt.Println("example_timeout")
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

func main() {
	exmaple_goroutine()
	exmaple_channel()
	example_Buffered_Channels()
	exmaple_range_close()
	exmaple_select()
	example_timeout()
}

/*
runtime goroutine
runtime包中有几个处理goroutine的函数：
Goexit
退出当前执行的goroutine，但是defer函数还会继续调用
Gosched
让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
NumCPU
返回 CPU 核数量
NumGoroutine
返回正在执 行和排队的任务总数
GOMAXPROCS
用来设置可以运行的CPU核数
*/
