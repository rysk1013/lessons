// 構造体・メソッド・ポインタレシーバー・値レシーバー
package main

import "fmt"

type Vertex struct {
	// X, Y int
	// 小文字にすると他のパッケージでは使えない
	x, y int
}

func (v Vertex) Area() int {
	// return v.X * x.Y
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	// v.X = v.X * i
	v.x = v.x * i
	// v.Y = v.Y * i
	v.y = v.y * i
}

func Area(v Vertex) int {
	// return v.X * x.Y
	return v.x * v.y
}

func New(x, y int) *Vertex {
	return &Vertex{x ,y}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area())

	v1 := New(3, 4)
	v1.Scale(10)
	fmt.Println(v.Area())
}