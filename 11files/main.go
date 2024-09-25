package main

import (
	"files/account"
	"files/files"
	"files/output"
    "strings"
	//"files/files"
	"fmt"

	"github.com/fatih/color"
	//"strings"
	//"strconv"
)


var menuMap = map[string]func(*account.VaultWithDb){
		"1": crAcc,
		"2": findAcc,
		"3": delAcc,
}

func main() {

	menu()

}

func menu() {
	var userCh string
	vault := account.NewVault(files.NewJsonDb("data.json"))

	 
	for {
		fmt.Println("Выберите функцию:")
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")
		fmt.Scanln(&userCh)
		
		menuFunc := menuMap[userCh]
		if menuFunc == nil {
				break 
		}

		menuFunc(vault)
				//switch userCh {
		//case 1:
		//	crAcc(vault)
		//	continue
		//case 2:
		//	findAcc(vault)
		//	continue
		//case 3:
		//	delAcc(vault)
		//	continue
		//default:
		//	return
		}

	}



func crAcc(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintErr("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)

}

func findAcc(vault *account.VaultWithDb) {
	var findUrl string
	fmt.Println("Введите URL необходимого аккаунта")
	fmt.Scanln(&findUrl)

	res, err := vault.FindAcc(findUrl, func(acc account.Account, str string)bool {
			return strings.Contains(acc.Url, str)
	})
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(res)

}




func delAcc(vault *account.VaultWithDb) {
	url := promptData("Введите URL аккаунта для удаления")
	isDeleted := vault.DeleteAccByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	}
	output.PrintErr("Не найдено")
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
