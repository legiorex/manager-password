package main

import "fmt"

type account struct {
    login string
    password string
    url string
}

func main() {
    login := promptData("Введите логин")
    password := promptData("Введите пароль")
    url := promptData("Введите URL")
    
    accountNew := account{
        login: login,
        password: password,
        url: url,
    }
    
    printAccount(&accountNew)

}

func promptData(message string) string {
    fmt.Print(message + ": ")
    var result string
    fmt.Scanln(&result)
    return result
}

func printAccount(acc *account) {
    fmt.Println("Логин: ", acc.login)
    fmt.Println("Пароль: ", acc.password)
    fmt.Println("URL: ", acc.url)
}
