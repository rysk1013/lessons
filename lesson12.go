// ストラクト
package main

import "fmt"

type Vertex struct {
	X int
	Y int
	// X, Y int
	S string
}

func changeVertex(vv Vertex) {
	vv.X = 1000
}

func changeVertex2(vv *Vertex) {
	vv.X = 1000
}

func main() {
	v6 := Vertex{1, 2, "test"}
	changeVertex(v6)
	fmt.Println(v6)

	v7 := &Vertex{1, 2, "test"}
	changeVertex2(v7)
	fmt.Println(*v7)

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "text"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	// varで宣言すると初期値が返ってくる
	// make(),map()はnilが返ってくる
	var v5 Vertex
	fmt.Println(v5)
}