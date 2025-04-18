package main

import (
	"fmt"

	"github.com/SANEKNAYMCHIK/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works!", st)
}
