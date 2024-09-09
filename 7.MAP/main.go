package main

import (
	"fmt"
	"os"
)

/*
	 Создать приложение, которое сначала выдаёт меню:
		1) Посмотреть закладки
		2) Добавить закладку
		3) Удалить закладку
		4) Выход

При 1 - выводит закладки
При 2 - 2 поля для ввода названия и адреса и после добавления
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/
type stringMap = map[string]string

var marks = make(stringMap, 100)

func main() {
	fmt.Println("Приложение для закладок")
	for {
		c := getUCh()

		switch c {
		case 1:
			listMarks(marks)
		case 2:
			addMark(marks)
			listMarks(marks)
		case 3:
			delMark(marks)
			listMarks(marks)
		case 4:
			os.Exit(0)

		default:
			fmt.Println("Неккоректный ввод")
			continue
		}

	}
}

func getUCh() int {
	var uCh int
	fmt.Print("1: Посмотреть закладки\n",
		"2: Добавить закладку\n",
		"3: Удалить закладку\n",
		"4: Выход\n")

	fmt.Scanln(&uCh)

	return uCh

}

func listMarks(marks stringMap) {
	fmt.Printf("\n\n__ Список ваших закладок __\n")
	if len(marks) == 0 {
		fmt.Println("Список ваших закладок пуст")
	}

	for k, v := range marks {

		if k != "" {
			fmt.Printf("%s: %s \n\n\n", k, v)

		}

	}
}

func addMark(m stringMap) {
	var key string
	var url string
	fmt.Println("__ Добавление заладки __")
	fmt.Println("Ввыедите название закладки")
	fmt.Scanln(&key)
	fmt.Println("Введите url сайта")
	fmt.Scanln(&url)

	m[key] = url

}

func delMark(m stringMap) {
	var key string
	fmt.Println("__ Удаление заладки __")
	fmt.Println("Ввыедите название закладки")
	fmt.Scanln(&key)

	delete(m, key)

}
