package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения
	if (strings.LastIndex(filePth, ".") < strings.LastIndex(filePth, "/")+1) || strings.LastIndex(filePth, ".") == -1 {
		fileName = filePth[strings.LastIndex(filePth, "/")+1 : len([]rune(filePth))]
		fileExt = ""
	} else {
		fileName = filePth[strings.LastIndex(filePth, "/")+1 : len([]rune(filePth))]
		fileExt = filePth[strings.LastIndex(filePth, "."):len([]rune(filePth))]
	}
	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)

	fmt.Println("Проверка")
	fmt.Printf("filename: %s\n", filepath.Base(filePth))
	fmt.Printf("extension: %s\n", filepath.Ext(filePth))
}
