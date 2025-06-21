package main

import (
	"fmt"

	"github.com/legiorex/manager-password/account"
	"github.com/legiorex/manager-password/files"
)

func main() {

	files.WriteFile("Hello", "text.txt")

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	// fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.PrintAccount()
	files.ReadFile()

}

func promptData(message string) string {
	fmt.Print(message + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}
