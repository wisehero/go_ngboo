package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a: ", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	fmt.Printf("a: %d b: %d f:%f\n", a, b, f) // 자동 개행 제공 X이기 때문에 끝에 \n 필요
	
}
