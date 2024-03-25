package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Вставьте ваш код здесь

	mp := map[rune]int{}
	n := 0
	for _, s := range text {
		if unicode.IsLetter(s) {
			mp[unicode.ToLower(s)]++
			n++
		}
	}

	for key, val := range mp {
		fmt.Printf("%c ", key)
		fmt.Printf("%0.1f\n", float64(val)/float64(n)*100)
	}
}
