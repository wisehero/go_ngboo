package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough // 되도록 사용하지 않는 것을 권장한다.
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	default:
		fmt.Println("a > 5")
	}
}
