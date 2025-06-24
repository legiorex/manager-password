package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/legiorex/manager-password/account"
)

var FILE_NAME = "pass.json"

func main() {

	vault := account.NewVault()

menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)

		case 2:
			searchAccount(vault)

		case 3:
			deleteAccount()

		default:
			break menu
		}
	}

}

func getMenu() int {

	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)

	green.Println("Создать аккаунт: 1")
	blue.Println("Найти аккаунт: 2")
	red.Println("Удалить аккаунт: 3")
	yellow.Println("Выход: 4")

	var variant int
	fmt.Scanln(&variant)

	return variant

}

func searchAccount(vault *account.Vault) {

	url := promptData("Введите URL")

	searchResult := vault.SearchAccountByUrl(url)

	if len(searchResult) == 0 {
		color.Cyan("Пароли не найдены")
	}

	for _, acc := range searchResult {
		acc.PrintAccount()
	}

}

func deleteAccount() {

	fmt.Println("deleteAccount")

}

func createAccount(vault *account.Vault) {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData(message string) string {
	fmt.Print(message + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}
