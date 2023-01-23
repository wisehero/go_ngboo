package main

import "fmt"

// 포인터를 사용하는 이유
// 1. 변수 대입이나 함수 인수 전달은 많은 메모리 공간을 차지
// 2. 다른 공간으로 복사되기 때문에 변경 사항이 적용되지도 않는다.

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
