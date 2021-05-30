package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("lesson9.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))
	//":="は一つでもイニシャライズするものがあれが再度宣言することができる

	//返り値が１つの場合は"="で上書きする
	//":="を使うとエラーになる
	err = os.Chdir("test")
	if err != nil {
		fmt.Println("Error!")
	}

	//返り値が１つの場合は短縮できる
	if err = os.Chdir("test"); err != nil {
		fmt.Println("Error!")
	}
}
