package main

import "fmt"

// Divide
// a int, b int를 아래처럼 줄일 수 있다.
// 함수 이름 시작이 대문자일 경우 패키지 외부 노출
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}
func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
