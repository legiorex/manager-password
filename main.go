package main

import (
	"fmt"

	"github.com/legiorex/manager-password/account"
	"github.com/legiorex/manager-password/files"
)

var FILE_NAME = "pass.json"

func main() {

	createAccount()

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

	file, err := myAccount.ToBytes()

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
