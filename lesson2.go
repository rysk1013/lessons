package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//数値型
	var (
		u8 uint8     	= 255
		i8 int8       = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v", u8, u8)
	
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)
	// 記述方法はターミナルでgofmt ファイル名
	// コードを書き換える場合はgofmt -w ファイル名
	

	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	fmt.Println(1 << 0) //0001 0001
	fmt.Println(1 << 1) //0001 0010
	fmt.Println(1 << 2) //0001 0100
	fmt.Println(1 << 3) //0001 1000

	//文字列型
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))

	fmt.Println(strings.Contains(s, "World"))

	//論理値型
	t, f := true, false
	fmt.Printf("%T, %v\n", t, t)
	fmt.Printf("%T, %v\n", f, f)

	//型変換
	var xyz int = 1
	xx := float64(xyz)
	fmt.Printf("%T, %v, %f\n", xx, xx, xx)

	var stu string = "14"
	i, _ := strconv.Atoi(stu)
	fmt.Printf("%T, %v", i, i)
}