package main

import "fmt"

func canIDrink(age int) bool {
	// if안에 변수 생성해서 사용 가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
func main() {
	fmt.Println(canIDrink(16))
}
