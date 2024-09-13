package storage

import (
	"bin/bins"
	"bin/file"
	"encoding/json"
	"fmt"
)

type Storage struct {
	Bins []bins.Bin
}

func (s *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewStorage() *Storage {
	f, err := file.Rfile("storage.json")
	if err != nil {
		fmt.Println(err)
		return &Storage{
			Bins: []bins.Bin{},
		}
	}

	var st Storage
	err = json.Unmarshal(f, &st)
	if err != nil {
		fmt.Println("не удалось запарсить JSON")
	}
	return &st
}

func (s *Storage) AddBin(bin bins.Bin) {
	s.Bins = append(s.Bins, bin)
	data, err := s.ToBytes()
	if err != nil {
		fmt.Println("Не удалось добавить элемент")
	}
	file.Wfile(data, "storage.json")
}
