package gor

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(n int) {
	fmt.Printf("hello index %v\n", n)
	defer wg.Done()
}

func Run() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func Run1() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func recv(c chan int) {
	// 通道中没有值时会阻塞
	ret := <-c
	fmt.Println("接收成功", ret)
}

func f2(ch chan int) {
	fmt.Println("通道容量", cap(ch))
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道关闭")
			wg.Done()
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func f3(ch chan int) {
	for v := range ch {
		fmt.Printf("v:%#v\n", v)
	}
	wg.Done()
}

func Run2() {
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int
	fmt.Println(ch1, ch2, ch3)
	// 无缓冲通道也叫同步通道
	ch4 := make(chan int)
	go recv(ch4)
	ch4 <- 10
	fmt.Println("发送成功")

	// 有缓冲通道 通道容量为1 当1没有接收者时会阻塞
	// 使用cap函数获取通道容量
	// 使用len获取通道内元素数量
	ch5 := make(chan int, 1)
	ch5 <- 10
	fmt.Println("发送成功")

	// 通道关闭后不能再发送值
	// 通道关闭后可以继续接收值
	ch6 := make(chan int, 2)
	ch6 <- 1
	ch6 <- 2
	close(ch6)
	wg.Add(1)
	// go f2(ch6)
	go f3(ch6)
	wg.Wait()

}

func Producer() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		// 任务完成以后关闭通道
		close(ch)
	}()
	return ch
}

func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func Run3() {
	ch := Producer()
	ret := Consumer(ch)
	fmt.Println(ret)
}

func Producer2() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		// 任务完成以后关闭通道
		close(ch)
	}()
	return ch
}

func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func Run4() {
	ch := Producer2()
	ret := Consumer2(ch)
	fmt.Println(ret)
}

func Run5() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
