package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func (i *Item) getItem() string {
	if i == nil {
		return " "
	}
	return fmt.Sprintf("Имя: %s\nURL: %s\nТеги: %s\nДата: %s\n", i.Name, i.Link, i.Tags, i.Date.Format(time.DateTime))
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода из приложения нажмите Esc")

	items := make([]*Item, 0)

	// Подключаем отслеживание нажатия клавиш
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}

OuterLoop:
	for {
		fmt.Print("Введите команду буквой (a)ppend, (l)ist, (r)emove>")

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		switch char {
		case 'a':
			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}

			// Напишите свой код здесь
			items = append(items, &Item{args[0], time.Now(), args[1], args[2]})
		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			fmt.Printf("Количество url: %d\n", len(items))
			for _, item := range items {
				fmt.Println(item.getItem())
			}
		case 'r':
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			name := strings.TrimSpace(text)

			// Напишите свой код здесь
			removed := false
			for i, item := range items {
				if item.Name == name {
					items = append(items[:i], items[i+1:]...)
					removed = true
					break
				}
			}

			if removed {
				fmt.Printf("Ссылка %s удалена\n", name)
			} else {
				fmt.Printf("Ссылка %s не найдена\n", name)
			}
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				fmt.Println("Выход")
				return
			}
		}
	}
}
