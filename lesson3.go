package main

import "fmt"

func main() {
	//配列
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// var b [2]int = [2]int{100, 200}
	// fmt.Println(b)
	//配列はサイズを変更することができない！
	//b = append(b, 300) ←３つ目の値を加えようとするとエラーになる
	//値を増やしたい場合はスライスを使う

	//スライス
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	var  board =  [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
	n = append(n, 100, 200, 300,400)
	fmt.Println(n)

	//スライスのmakeとcap
	n1 := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n1), cap(n1), n1)
	n1 = append(n1, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n1), cap(n1), n1)
	n1 = append(n1, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n1), cap(n1), n1)

	a1 := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a1), cap(a1), a1)

	b1 := make([]int, 0)
	var c1 []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b1), cap(b1), b1)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c1), cap(c1), c1)

	//c2とc3の違いを理解する
	c2 := make([]int, 5) 
	for i := 0; i < 5; i++ {
		c2 = append(c2, i)
		fmt.Println(c2)
	}
	fmt.Println(c2)

	c3 := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c3 = append(c3, i)
		fmt.Println(c3)
	}
	fmt.Println(c3)

	//map
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	//記述されていないものを指定したとき
	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	//初期化せずに値を代入するとエラーになる
	// var m3 map[string]int
	// m3["tv"] = 3000
	// fmt.Println(m3)

	//varで宣言するとnilが返ってくる
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}