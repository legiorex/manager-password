package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/legiorex/manager-password/files"
)

type Vault struct {
	Accounts  []AccountWithTimeStamp `json:"accounts"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func NewVault(db *files.JsonDb) *VaultWithDb {

	file, err := db.Read()

	if err != nil {

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamp{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red(err.Error())

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamp{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    *db,
	}

}

func (vault *VaultWithDb) AddAccount(acc AccountWithTimeStamp) error {
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

func (vault *VaultWithDb) SearchAccountByUrl(url string) []AccountWithTimeStamp {

	var searchResult []AccountWithTimeStamp

	for _, acc := range vault.Accounts {

		isSearch := strings.Contains(acc.Url, url)

		if isSearch {
			searchResult = append(searchResult, acc)
		}

	}

	return searchResult

}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {

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

func (vault *VaultWithDb) save() error {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()

	if err != nil {
		return err
	}
	vault.db.Write(data)
	return nil
}
