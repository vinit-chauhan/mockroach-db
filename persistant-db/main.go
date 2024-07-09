package main

import (
	"fmt"
	"persistant-db/pkg"
)

func main() {
	instance := new(pkg.JsonDB)
	db, err := instance.New("./")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	for _, val := range pkg.GetEmployeesSlice() {
		db.Write(val)
	}

}
