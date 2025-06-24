package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/legiorex/manager-password/files"
)

var FILE_NAME = "pass.json"

type Vault struct {
	Accounts  []AccountWithTimeStamp `json:"accounts"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

func NewVault() *Vault {

	file, err := files.ReadFile(FILE_NAME)

	if err != nil {

		return &Vault{
			Accounts:  []AccountWithTimeStamp{},
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

func (vault *Vault) AddAccount(acc AccountWithTimeStamp) error {
	vault.Accounts = append(vault.Accounts, acc)

	err := vault.save()

	if err != nil {
		return err
	}

	return nil
}

func (vault *Vault) ToBytes() ([]byte, error) {

	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) SearchAccountByUrl(url string) []AccountWithTimeStamp {

	var searchResult []AccountWithTimeStamp

	for _, acc := range vault.Accounts {

		isSearch := strings.Contains(acc.Url, url)

		if isSearch {
			searchResult = append(searchResult, acc)
		}

	}

	return searchResult

}

func (vault *Vault) DeleteAccountByUrl(url string) bool {

	var searchResult []AccountWithTimeStamp

	isDelete := false

	for _, acc := range vault.Accounts {

		isSearch := strings.Contains(acc.Url, url)

		if !isSearch {
			searchResult = append(searchResult, acc)
			continue
		}
		isDelete = true

	}

	vault.Accounts = searchResult

	vault.save()
	return isDelete

}

func (vault *Vault) save() error {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()

	if err != nil {
		return err
	}
	files.WriteFile(data, FILE_NAME)
	return nil
}
