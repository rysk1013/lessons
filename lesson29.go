package main

import (
	"fmt"
	"sort"
	"regexp"
	"time"
)

func main() {
	// 標準パッケージ
	// <time>
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// <regex>(正規表現)
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch(("/view/test"))
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch(("/edit/test"))
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch(("/save/test"))
	fmt.Println(fss, fss[0], fss[1], fss[2])

	// <sort>
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool {return p[i].Name < p[j].Name})
	sort.Slice(p, func(i, j int) bool {return p[i].Age < p[j].Age})
	fmt.Println(i, s, p)

	// <iota>
	const (
		c1 = iota
		c2
		c3
	)

	const (
		_ = iota
		KB int = 1 << (10 * iota)
		MB
		GB
	)
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}