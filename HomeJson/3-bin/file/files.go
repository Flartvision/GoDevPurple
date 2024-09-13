package file

import (
	"fmt"
	"os"
)

func Rfile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return nil, err

	}

	return data, nil

}

func Wfile(content []byte, name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.Write(content)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")

}
