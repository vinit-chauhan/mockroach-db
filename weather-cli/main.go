package main

import (
	"os"
	app "weather-cli/pkg"
)

func main() {
	location := "Windsor"

	if len(os.Args) > 1 {
		location = os.Args[1]
	}

	app.Run(location)
}
