package files

import (
	"fmt"
	"os"
)

func WriteF(content []byte, name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = f.Write(content)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")

}

func ReadF(name string) ([]byte, error) {
	d, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)

	}
	return d, err

}
