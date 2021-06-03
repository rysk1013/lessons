package main

import "fmt"

func ask(num int ,q string, p *int) {
	fmt.Printf("[Q%d]Type the following: %v\n", num, q)
	var input string
	fmt.Scan(&input)

	if q == input {
		fmt.Println("Good!")
		*p += 10
	} else {
		fmt.Println("Bad...")
	}
}

func main() {
	totalScore := 0
	ask(1, "html", &totalScore)
	ask(2, "css", &totalScore)
	ask(3, "golang", &totalScore)
	ask(4, "ruby", &totalScore)

	fmt.Println("スコア:", totalScore )
}