package main

import (
	"fmt"
)

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccountWithTimeStamp(login, password, url)

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
