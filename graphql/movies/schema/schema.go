package schema

import (
	"os"
)

func String() (string, error) {
	file, err := os.ReadFile("./movies/schema/schema.graphql")
	if err != nil {
		return "", err
	}
	return string(file), nil
}
