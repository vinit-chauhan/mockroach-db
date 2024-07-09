package pkg

import "errors"

var dir string = "./"

func Init() MockroachDB {
	return MockroachDB{}
}

func (m *MockroachDB) New(path string) (*MockroachDB, error) {
	return nil, errors.New("Not Implemented yet!!!")
}
