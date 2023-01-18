package main

import "fmt"

/*
func Add(a int, b int) int {
	// return a + b
}
*/

// 대문자는 패키지 외부로 노출하는 함수. public
func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}
