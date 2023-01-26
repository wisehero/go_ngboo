package main

import "fmt"

func main() {

	// var str string = "Hello World"
	// str = "How are you?" ->  전체 바꾸기는 가능
	// str[2] = 'a' -> Error!
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
