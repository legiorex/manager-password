package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/legiorex/manager-password/account"
	"github.com/legiorex/manager-password/files"
	"github.com/legiorex/manager-password/output"
)

var FILE_NAME = "pass.json"

func main() {

	vault := account.NewVault(files.NewJsonDb(FILE_NAME))

menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)

		case 2:
			searchAccount(vault)

		case 3:
			deleteAccount(vault)

		default:
			break menu
		}
	}

}

func getMenu() int {

	// promptData2([]string{"Hello", "Buy"})

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

func searchAccount(vault *account.VaultWithDb) {

	url := promptData("Введите URL")

	searchResult := vault.SearchAccountByUrl(url)

	if len(searchResult) == 0 {
		color.Cyan("Пароли не найдены")
	}

	for _, acc := range searchResult {
		acc.PrintAccount()
	}

}

func deleteAccount(vault *account.VaultWithDb) {

	url := promptData("Введите URL")

	isDelete := vault.DeleteAccountByUrl(url)

	if isDelete {
		color.Green("Аккаунт успешно удален")
	} else {

		output.PrintError("Ошибка удаления аккаунта")

	}

}

func createAccount(vault *account.VaultWithDb) {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		output.PrintError(err)
		return
	}

	err = vault.AddAccount(*myAccount)
	if err != nil {
		output.PrintError("Ошибка при сохранении аккаунта")
	}
	color.Green("Запись успешна")
}

func promptData(message string) string {
	fmt.Print(message + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}

// func promptData2[T any](values []T) string {

// 	for i, item := range values {

// 		if i == len(values)-1 {
// 			fmt.Printf("%v :", item)
// 		} else {

// 			fmt.Println(item)
// 		}

// 	}

// 	var result string
// 	fmt.Scanln(&result)
// 	return result
// }
