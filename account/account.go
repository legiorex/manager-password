package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

type AccountWithTimeStamp struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Account
}

func (acc *Account) PrintAccount() {
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Логин: ", acc.Login)
	c.Println("Пароль: ", acc.Password)
	c.Println("URL: ", acc.Url)
}

func (acc *Account) generatePassword(n int) {

	// создаем 32-байтное семя для инициализации генератора
	seed := [32]byte{'s', 'o', 'm', 'e', 'k', 'e', 'y', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'a', 'b', 'c', 'd', 'e', 'f'}

	// инициализируем новый генератор ChaCha8
	rand.NewChaCha8(seed)

	// start 65
	// finish 90

	min := 65
	max := 90

	var strBuilder strings.Builder

	for range n {
		genRune := rand.IntN(max-min+1) + min

		strBuilder.WriteRune(rune(genRune))

	}

	resultPassword := strBuilder.String()

	acc.Password = resultPassword
}

// func newAccount(login, password, urlUser string) (*account, error) {

// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}

// 	_, errUrl := url.ParseRequestURI(urlUser)

// 	if errUrl != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}

// 	acc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlUser,
// 	}

// 	if acc.password == "" {
// 		acc.generatePassword(12)
// 	}

// 	return acc, nil
// }

func NewAccountWithTimeStamp(login, password, urlUser string) (*AccountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, errUrl := url.ParseRequestURI(urlUser)

	if errUrl != nil {
		return nil, errors.New("INVALID_URL")
	}

	acc := &AccountWithTimeStamp{

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Account: Account{
			Login:    login,
			Password: password,
			Url:      urlUser,
		},
	}

	if acc.Password == "" {
		acc.generatePassword(12)
	}

	return acc, nil
}
