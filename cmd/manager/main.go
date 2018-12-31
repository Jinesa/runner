package main

import (
	"github.com/deissh/lambda/pkg/manager"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	m, _ := manager.Create()

	active, err := m.GetAll()
	if err != nil {
		log.Panic(err)
	}

	for _, container := range active{
		log.Println(container.ID)
	}
}