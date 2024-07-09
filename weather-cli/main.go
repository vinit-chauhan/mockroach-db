package main

import (
	"fmt"
	"time"
)

func greet() string {
	hour := time.Now().Local().Hour()

	if hour > 4 && hour < 12 {
		return "Morning"
	} else if hour >= 12 && hour < 5 {
		return "Afternoon"
	} else if hour >= 5 && hour < 9 {
		return "Evening"
	} else {
		return "Night"
	}
}

func main() {
	fmt.Printf("Good %s!!!\n", greet())

}
