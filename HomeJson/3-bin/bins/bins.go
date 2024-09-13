package bins

import (
	"encoding/json"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	data []Bin
}

func (b *Bin) ToBytes() ([]byte, error) {
	file, err := json.Marshal(b)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewBin(id string, name string, p bool) *Bin {

	return &Bin{
		Id:        id,
		Private:   p,
		CreatedAt: time.Now(),
		Name:      name,
	}

}
