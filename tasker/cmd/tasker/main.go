package main

import (
	"fmt"

	"github.com/vinit-chauhan/tasker/internal/db"
)

func main() {
	path := db.SetupPath()

	fmt.Println(path)
}
