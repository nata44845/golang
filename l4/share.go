package main

import (
	"fmt"
)

type mySlice []string

func (m mySlice) String() string {
	return fmt.Sprintf("%#v", m)
}

func main() {
	m := mySlice{"hello", "world"}
	fmt.Println(m)
}
