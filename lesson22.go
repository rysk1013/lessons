// Channel・Buffered Channel
package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	// Channel
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // 15, 15...
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <- c
	fmt.Println(x)
	y := <- c
	fmt.Println(y)

	// Buffered Channel
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch) 
	
	// rangeでchannelの値を取り出すときはclose()で閉じる
	for c1 := range ch {
		fmt.Println(c1)
	}

	// channel/close/range
	s1 := []int{1, 2, 3, 4, 5}
	c1 := make(chan int)
	go goroutine3(s1, c1)
	for i := range c1 {
		fmt.Println(i)
	}
}