package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Println(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str) // 특수문자까지 모두 출력해버림
}
