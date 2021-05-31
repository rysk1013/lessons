// ポインタ
package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n) //100が入っているメモリーを示している

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	var  n1 int =  100
	one(&n)
	fmt.Println(n1)

	var p1 *int = new(int) //new()を使うとメモリが確保される
	fmt.Println(p1)

	var p2 *int // 変数宣言だけではメモリは確保されずnilとなる
	fmt.Println(p2)

	// make()とnew()の違い
	// new()はポインタが作成される
	slice := make([]int, 0)
	fmt.Printf("%T\n", slice)

	m := make(map[string] int)
	fmt.Printf("%T\n", m)

	var p4 *int = new(int)
	fmt.Printf("%T\n", p4)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}