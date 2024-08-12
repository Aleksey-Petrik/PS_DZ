package main

import (
	"log"
	"main/storage"
)

func main() {

	s, err := storage.NewStorage("storage.json")
	if err != nil {
		log.Println(err)
	} else {
		/*		b1 := bins.Bin{ID: "1", Private: false, CreateAt: time.Now(), Name: "Test 1"}
				s.AddBin(b1)
				b2 := bins.Bin{ID: "2", Private: false, CreateAt: time.Now(), Name: "Test 2"}
				s.AddBin(b2)*/

		s.Print()
	}
}
