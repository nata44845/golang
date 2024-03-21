package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) bool {
	i := 2
	for i < n {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

func main() {
	// В данном примере мы сознатель опускаем проверку типов.
	// Если передать в качестве ввода значение типа строка или булево значение, то программа аварийно завершится
	// Пакет log предоставляет нам базовые функции логирования.
	// log.Fatal печатает на экране переданный текст ошибки и звершается
	var value int
	fmt.Println("Введите число")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal(err)
	}

	if isPrime(value) {
		fmt.Println("Число", value, "простое")
	} else {
		fmt.Println("Число", value, "не простое")
	}
}
