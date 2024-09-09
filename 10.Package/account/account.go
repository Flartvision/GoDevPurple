package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
	//"strings"
	//"strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVXYZ123456-*")

type account struct {
	login    string
	password string
	url      string
}

type accWithTimestamp struct {
	account
	createAt time.Time
	upAt     time.Time
}

func (acc *account) OutPass() {
	color.Cyan(acc.login)

}

func (acc *account) genPass(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccWithTimeStamp(login, password, urlString string) (*accWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("invalid_login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	newAcc := &accWithTimestamp{
		createAt: time.Now(),
		upAt:     time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
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
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		newAcc.genPass(12)
	}
	return newAcc, nil
}
