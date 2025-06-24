package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []AccountWithTimeStamp `json:"accounts"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts:  []AccountWithTimeStamp{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc AccountWithTimeStamp) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) {

	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}
	return file, nil
}
