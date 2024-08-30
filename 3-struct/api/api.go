package api

import (
	"log"
	"main/config"
)

func NewApi(cfg config.Config) {
	log.Println("KEY - ", cfg.Key)
}
