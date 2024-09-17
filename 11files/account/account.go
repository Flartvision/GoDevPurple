package account

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/url"

	"github.com/fatih/color"
	//"strings"
	//"strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVXYZ123456-*")

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func (acc *Account) OutPass() {
	color.Cyan(acc.Login)

}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) genPass(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("invalid_login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	newAcc := &Account{
		Url:      urlString,
		Login:    login,
		Password: password,
	}
	if password == "" {
		newAcc.genPass(12)
	}
	return newAcc, nil
}
