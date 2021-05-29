//function

package main

import "fmt"

func add(x, y int) (int, int) {
	//fmt.Println("add function")
	return x + y, x - y
}

func calc(price, item int) (result int) {
	result = price * item
	return result //resultは省略してもok
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := calc(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int)  {
		fmt.Println("inner func", x)
	}(1)
}