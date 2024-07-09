package main

import (
	"fmt"
	"mockroach-db/pkg"
)

func main() {
	instance := new(pkg.MockroachDB)
	db, err := instance.New("./")
	if err != nil {
		fmt.Println("ERROR: %w", err)
	}

	for _, val := range pkg.GetEmployeesSlice() {
		db.Write(val)
	}

}
