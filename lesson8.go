package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
//defer(遅延実行)
defer fmt.Println("world")
foo()
fmt.Println("hello")

fmt.Println("run")    	//実行される順番１
defer fmt.Println(1)		//実行される順番５
defer fmt.Println(2)		//実行される順番４
defer fmt.Println(3)		//実行される順番３
fmt.Println("success")	//実行される順番２

// deferが使われる例
file, _ := os.Open("lesson7.go")
defer file.Close()
data := make([]byte, 100)
file.Read(data)
fmt.Println(string(data))
//ファイルをオープンしたときは必ずクローズしないといけない
//deferを使うことでクローズし忘れることがなくなる
}