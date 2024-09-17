package files

import (
	"files/output"
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}

}

func (db *JsonDb) Write(content []byte) {
	f, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = f.Write(content)
	defer f.Close()
	if err != nil {
		output.PrintErr(err)
		return
	}
	fmt.Println("Запись успешна")

}

func (db *JsonDb) Read() ([]byte, error) {
	d, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err

	}
	return d, nil

}
