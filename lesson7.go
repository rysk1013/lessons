package main

import (
	"fmt"
	"time"
)

// if文で使用
func by2(num int) string {
	if num % 2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

//switchで使用
func getOsName() string {
	return "aaa"
}

func main() {
	// if文

	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result) //resultを使ってもエラーにならない


	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	// fmt.Println(result2) //result2を使おうとするとエラーになる

	num := 6
	if num % 2 == 0 {
		fmt.Println("by 2")
	} else if num % 3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}

	// for文
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for; sum < 10; {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// range
	l := []string{"python", "golang", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	// switch
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Wondows!")
	default:
		fmt.Println("Default")
	}

	//switchでのみ変数を使う場合の記述方法
	switch os1 := getOsName(); os1 {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Wondows!")
	default:
		fmt.Println("Default", os1)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	}
}