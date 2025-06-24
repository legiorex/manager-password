package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {

	data, err := os.ReadFile(db.fileName)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.fileName)

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
