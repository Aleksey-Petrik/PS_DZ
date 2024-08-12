package storage

import (
	"encoding/json"
	"log"
	"main/bins"
	"main/files"
	"os"
	"time"
)

type Storage struct {
	Bins     []bins.Bin `json:"bins"`
	UpdateAt time.Time  `json:"createAt"`
	FilePath string
}

func NewStorage(filePath string) (*Storage, error) {
	bins, err := read(filePath)

	if err != nil {
		return nil, err
	}

	return &Storage{Bins: bins, UpdateAt: time.Now(), FilePath: filePath}, nil
}

func (s *Storage) save() error {
	data, err := json.Marshal(s.Bins)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.FilePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) AddBin(bin bins.Bin) {
	s.Bins = append(s.Bins, bin)
	s.save()
}

func (s *Storage) Print() {
	log.Println(s.Bins)
}

func read(filePath string) ([]bins.Bin, error) {
	data, err := files.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	bins, err := dataToBians(data)
	if err != nil {
		return nil, err
	}
	return bins, nil
}

func dataToBians(data []byte) ([]bins.Bin, error) {
	var bins []bins.Bin
	err := json.Unmarshal(data, &bins)
	if err != nil {
		return nil, nil
	}
	return bins, nil
}
