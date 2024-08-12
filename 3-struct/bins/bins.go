package bins

import "time"

type Bin struct {
	ID       string    `json:"id"`
	Private  bool      `json:"private"`
	CreateAt time.Time `json:"createAt"`
	Name     string    `json:"name"`
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		ID:       id,
		Private:  private,
		Name:     name,
		CreateAt: time.Now(),
	}
}
