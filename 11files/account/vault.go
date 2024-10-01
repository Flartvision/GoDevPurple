package account

import (
	"encoding/json"
	"errors"
	"files/encrypter"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {

	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red(err.Error())
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}

}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()

}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) FindAcc(str string, checker func(Account, string) bool) ([]Account, error) {
	var finder []Account
	for _, v := range vault.Accounts {
		isMatched := checker(v, str)
		if isMatched {
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

func (vault *VaultWithDb) DeleteAccByUrl(url string) bool {
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
	vault.Accounts = accounts
	vault.save()
	return isDeleted

}

func (v *VaultWithDb) save() {
	v.UpdatedAt = time.Now()
	data, err := v.Vault.ToBytes()
	encData := v.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	v.db.Write(encData)
}
