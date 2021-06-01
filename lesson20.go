package main

import "fmt"

// 演習１
type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y 
}

// 演習２
type Vertex2 struct {
	X, Y int
}

func (v Vertex2) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}

func main() {
// 演習１
	v := Vertex{3, 4}
	fmt.Println(v.Plus())

// 演習２
	v2 := Vertex2{3, 4}
	fmt.Println(v2)
}