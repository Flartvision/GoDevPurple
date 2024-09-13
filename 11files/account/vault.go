package account

import (
	"encoding/json"
	"errors"
	"files/files"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadF("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault

}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать JSON")
	}
	files.WriteF(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccByURL(url string) ([]Account, error) {
	var finder []Account
	for _, v := range vault.Accounts {
		if strings.Contains(v.Url, url) {
			finder = append(finder, v)
			fmt.Println(v.Login, v.Password)
		}
		continue
	}
	if finder == nil {
		return nil, errors.New("аккаунтов не найдено")
	}
	return finder, nil
}

func (vault *Vault) DeleteAccByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if !isMatched {
			accounts = append(accounts, acc)
			continue

		}
		isDeleted = true
	}

	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать JSON")
	}
	files.WriteF(data, "data.json")
	return isDeleted

}
