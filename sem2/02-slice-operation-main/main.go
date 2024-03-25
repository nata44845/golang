package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"
)

// Напишите функцию фильтрации слайса. Отфильтруйте слайс arr1 так, чтобы он содержал только не четные числа
// То есть, например arr1 = [0, 2, 3, 1, 5, 4] после фильтрации [3, 1, 5]
// Не используйте готовые функции из пакета slices
func filter(arr1 []int) []int {
	newarr := make([]int, 0, len(arr1))

	for _, s := range arr1 {
		if s%2 == 1 {
			newarr = append(newarr, s)
		}
	}

	return newarr
}

// Напишите функцию вставки элемента в слайс на позицию
// То есть, например arr1 = [0, 2, 3, 1, 5, 4] pos = 1, value = 4. Результат [0, 4, 2, 3, 1, 5, 4]
// Не используйте готовые функции из пакета slices
func insert(pos, value int, arr1 []int) []int {
	return append(arr1[:pos], append([]int{value}, arr1[pos:]...)...)
	// newarr := make([]int, 0, len(arr1))
	// newarr = append(newarr, arr1[:pos]...)
	// newarr = append(newarr, value)
	// newarr = append(newarr, arr1[pos:]...)
	// return newarr
}

func main() {
	// Этот код нужен для ввода массива из стандартного ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите элементы массива через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем строку на массив строк
	strValues := strings.Split(input, " ")

	// Преобразуем строки в числа и заполняем массив
	var arr []int
	for _, str := range strValues {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		arr = append(arr, val)
	}

	var pos int
	fmt.Println("Введите позицию для вставки")
	pos1, _ := reader.ReadString('\n')
	pos1 = strings.TrimSpace(pos1)
	pos, _ = strconv.Atoi(pos1)
	// if _, err := fmt.Fscan(os.Stdin, pos); err != nil {
	// 	log.Fatal("Неправильное значение")
	// }
	// if pos < 0 || pos > len(arr)-1 {
	// 	log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	// }

	var value int
	fmt.Println("Введите значение для вставки")
	value1, _ := reader.ReadString('\n')
	value1 = strings.TrimSpace(value1)
	value, _ = strconv.Atoi(value1)
	// if _, err := fmt.Fscan(os.Stdin, value); err != nil {
	// 	log.Fatal("Неправильное значение")
	// }
	// if pos < 0 || pos > len(arr)-1 {
	// 	log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	// }

	// Скопируйте слайс arr в слайс arr1.
	// Это можно сделать через функцию Append или через функцию copy
	arr1 := make([]int, 0, len(arr))
	arr1 = append(arr1, arr[:]...)

	arr1 = insert(pos, value, arr1)

	arr1 = filter(arr1)

	fmt.Println(arr1)
}
