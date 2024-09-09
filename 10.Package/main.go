package main

import (
	"fmt"
	"package/account"
	//"strings"
	//"strconv"
)

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1, err := account.NewAccWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login")
		return
	}

	account1.OutPass()
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
