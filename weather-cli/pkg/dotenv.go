package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var vars = map[string]string{}

func GetVars() (map[string]string, error) {
	file, err := os.Open(".env")
	if err != nil {
		return nil, errors.New("error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		pair := strings.Split(line, "=")

		if len(pair) != 2 {
			return nil, errors.New(fmt.Sprintf("no value provided for %s", pair[0]))
		}

		vars[pair[0]] = pair[1]
	}

	return vars, nil
}
