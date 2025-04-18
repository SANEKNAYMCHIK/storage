package main

import (
	"fmt"
	"log"

	"github.com/SANEKNAYMCHIK/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it uploaded", file)
}
