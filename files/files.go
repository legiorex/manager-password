package files

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) {

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)

	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
	file.Close()

}
