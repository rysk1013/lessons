// 可変長引数
package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)

	//練習問題１
	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	//練習問題２
	//[5,6]

	//練習問題３
	m := map[string]int {
		"Mike": 20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}