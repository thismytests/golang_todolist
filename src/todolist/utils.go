package todolist

import (
	"log"
	"math/rand"
)

func generateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b[:5])
}