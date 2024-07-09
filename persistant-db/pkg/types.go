package pkg

import "encoding/json"

type JsonDB struct {
	allUsers []User
}

type User struct {
	Name    string      `json:"name"`
	Age     json.Number `json:"age"`
	Contact string      `json:"contact"`
	Company string      `json:"company"`
	Address Address     `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Pincode string `json:"pincode"`
}
