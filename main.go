package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) printAccount() {
	fmt.Println("Логин: ", acc.login)
	fmt.Println("Пароль: ", acc.password)
	fmt.Println("URL: ", acc.url)
}

func (acc *account) generatePassword(n int) {

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

	acc.password = resultPassword
}

func newAccount(login, password, urlUser string) (*account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, errUrl := url.ParseRequestURI(urlUser)

	if errUrl != nil {
		return nil, errors.New("INVALID_URL")
	}

	acc := &account{
		login:    login,
		password: password,
		url:      urlUser,
	}

	if acc.password == "" {
		acc.generatePassword(12)
	}

	return acc, nil
}

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccount(login, password, url)

	// fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.printAccount()

}

func promptData(message string) string {
	fmt.Print(message + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}
