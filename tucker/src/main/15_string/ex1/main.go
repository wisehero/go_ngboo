package main

import "fmt"

func main() {
	str1 := "Hello \t 'World' \n"
	str2 := `Go is "awesome"!\nGo is simple and\t 'powerful''` // 특수문자 동작 X

	fmt.Println(str1)
	fmt.Println(str2)
}
