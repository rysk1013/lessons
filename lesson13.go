package main

import "fmt"

func main() {
	// 演習１
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j) //10

	//演習２
	var a int = 100
	var b int = 200
	var p1 *int
	var p2 *int
	p1 = &a
	p2 = &b
	a = *p1 + *p2
	p2 = p1
	j = *p2 + a
	fmt.Println(j) //600
}