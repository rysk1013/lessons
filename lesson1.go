// 変数宣言・const

package main

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

var (
	i int = 1
	f64 float64 = 1.2
	s string = "text"
	t, f bool = true, false
)

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
}

func main() {
	fmt.Println(i, f64, s, t, f)

	foo()

	fmt.Println(Pi, Username, Password)
}