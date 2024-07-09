package pkg

import (
	"errors"
	"fmt"
)

var dir string = "./"

func Init() JsonDB {
	return JsonDB{}
}

func (m *JsonDB) New(path string) (*JsonDB, error) {
	return nil, errors.New("Not Implemented yet!!!")
}

func (d *JsonDB) Write(User) {
	fmt.Println("Added a user")
}
