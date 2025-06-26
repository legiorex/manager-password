package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/legiorex/manager-password/cryptography"
	"github.com/legiorex/manager-password/output"
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

// type Db interface {
// 	Read() ([]byte, error)
// 	Write([]byte)
// }

type Vault struct {
	Accounts  []AccountWithTimeStamp `json:"accounts"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db           Db
	cryptography cryptography.Cryptography
}

func NewVault(db Db, cryptography cryptography.Cryptography) *VaultWithDb {

	file, err := db.Read()

	if err != nil {

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamp{},
				UpdatedAt: time.Now(),
			},
			db:           db,
			cryptography: cryptography,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		output.PrintError(err)

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamp{},
				UpdatedAt: time.Now(),
			},
			db:           db,
			cryptography: cryptography,
		}
	}

	return &VaultWithDb{
		Vault:        vault,
		db:           db,
		cryptography: cryptography,
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

func (vault *VaultWithDb) SearchAccount(str string, checker func(*AccountWithTimeStamp, string) bool) []AccountWithTimeStamp {

	var searchResult []AccountWithTimeStamp

	for _, acc := range vault.Accounts {

		isSearch := checker(&acc, str)

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
