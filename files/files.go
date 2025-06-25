package files

import (
	"fmt"
	"os"

	"github.com/legiorex/manager-password/output"
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
		output.PrintError(err)
	}

	_, err = file.Write(content)

	if err != nil {
		file.Close()
		output.PrintError(err)

		return
	}
	fmt.Println("Запись успешна")
	file.Close()

}
