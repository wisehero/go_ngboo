package main

import "fmt"

func main() {
	str := "Hello World"
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	// string과 rune 슬라이스는 상호 변환이 가능하다.
	fmt.Println(string(runes))
}
