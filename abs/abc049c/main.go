package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	w1 := "eraser"
	w2 := "erase"
	w3 := "dreamer"
	w4 := "dream"
	// 反転
	s = reverse(s)
	s = strings.Replace(s, reverse(w1), "", -1)
	s = strings.Replace(s, reverse(w2), "", -1)
	s = strings.Replace(s, reverse(w3), "", -1)
	s = strings.Replace(s, reverse(w4), "", -1)

	var res string
	if s == "" {
		res = "YES"
	} else {
		res = "NO"
	}

	fmt.Println(res)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
