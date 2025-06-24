package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/legiorex/manager-password/account"
	"github.com/legiorex/manager-password/files"
)

var FILE_NAME = "pass.json"

func main() {
menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()

		case 2:
			searchAccount()

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

func searchAccount() {

	fmt.Println("searchAccount")

}

func deleteAccount() {

	fmt.Println("deleteAccount")

}

func createAccount() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)

	file, err := vault.ToBytes()

	if err != nil {
		fmt.Println(err)
		return
	}

	files.WriteFile(file, FILE_NAME)

	files.ReadFile(FILE_NAME)

}

func promptData(message string) string {
	fmt.Print(message + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}
