package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [100]int
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range arr {
		value := rnd.Intn(20)
		arr[i] = value
	}
	fmt.Println(arr)
	// Создайте слайс нужного размера
	slice := make([]int, 20, 20)

	// Заполните слайс количеством встречаемых чисел
	for _, val := range arr {
		slice[val]++
	}

	fmt.Println(slice)
}
