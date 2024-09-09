package main

import (
	"files/account"
	"files/files"

	//"files/files"
	"fmt"
	//"strings"
	//"strconv"
)

func main() {
	menu()

}

func menu() {
	var userCh int

	for {
		fmt.Println("Выберите функцию:")
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")
		fmt.Scanln(&userCh)

		switch userCh {
		case 1:
			crAcc()
			continue
		case 2:
			findAcc()
			continue
		case 3:
			delAcc()
			continue
		default:
			return
		}

	}

}

func crAcc() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login")
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteF(data, "data.json")

}

func findAcc() {
	fmt.Println("Заглушка для поиска")
}

func delAcc() {
	fmt.Println("Заглушка для удаления")
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
