package main

import "fmt"

func main() {
    login := promptData("Введите логин: ")
    password := promptData("Введите пароль: ")

}

func promptData(message string) string {
    fmt.Println(message)
    var result string
    fmt.Scanln(&result)
    return result
}
