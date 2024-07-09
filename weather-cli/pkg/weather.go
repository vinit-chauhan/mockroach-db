package internal

import (
	"fmt"
	"time"
)

func greet() string {
	hour := time.Now().Local().Hour()
	if hour > 4 && hour < 12 {
		return "Morning"
	} else if hour >= 12 && hour < 17 {
		return "Afternoon"
	} else if hour >= 17 && hour < 21 {
		return "Evening"
	} else {
		return "Night"
	}
}

func Run() {
	fmt.Printf("Good %s!!!\n", greet())
	env, err := GetVars()
	if err != nil {
		panic(fmt.Sprintf("PANIC: Unable to read env file: %s", err.Error()))
	}

	fmt.Println(env["API_KEY"])
}
