package main

import "fmt"

type account struct {
    login string
    password string
    url string
}

func main() {
    login := promptData("Введите логин: ")
    password := promptData("Введите пароль: ")
    url := promptData("Введите URL: ")

}

func promptData(message string) string {
    fmt.Println(message)
    var result string
    fmt.Scanln(&result)
    return result
}
