package main

import (
	"fmt"
	"math/rand/v2"
	"errors"
	"net/url"
	"time"
	//"strings"
	//"strconv"
)


type account struct{
	login string
	password string
	url string
}

type accWithTimestamp struct {
	account
	createAt time.Time
	upAt tume.Time
}

func (acc *account)outPass() {
	fmt.Println(acc.login, acc.password, acc.url)
}


func (acc *account)genPass(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccWithTimeStamp(login, password, urlString string) (*accWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("invalid_login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	newAcc := &accWithTimestamp{
		createAt: time.Now(),
		upAt time.Now(),
		account: account{
			url: urlString,
			login: login,
			password: password,
		}
}
	if password == "" {
		newAcc.genPass(12)
	}
	return newAcc, nil
}

func newAcc(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("invalid_login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	newAcc := &account{
		url: urlString,
		login: login,
		password: password,
}
	if password == "" {
		newAcc.genPass(12)
	}
	return newAcc, nil
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVXYZ123456-*")

func main() {
	delete windows && android
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1, err := newAccWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login")
		return 
	}
	
	account1.outPass()
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}





