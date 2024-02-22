package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

func (a *StructA) AAA(x int) int {
	return x * x
}

func (a *StructA) BBB(x int) string {
	return "X=" + strconv.Itoa(x)
}

type StructB struct {
}

func (a *StructB) AAA(x int) int {
	return x * 2
}

// BBB가 구현되지 않으면 StructB는 InterfaceA를 implement하지 않는다.
func (a *StructB) BBB(x int) string {
	return "X=" + strconv.Itoa(x)
}

func main() {
	var c InterfaceA = &StructA{}

	var d InterfaceA = &StructB{}

	fmt.Println(c.AAA(3))
	fmt.Println(c.BBB(3))

	fmt.Println(d.AAA(3))
}
