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

func (vault *Vault) AddAccount(acc AccountWithTimeStamp) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()

	if err != nil {
		color.Red(err.Error())
	}

	files.WriteFile(data, FILE_NAME)

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
