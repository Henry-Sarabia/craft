package craft

import (
	"errors"
	"math/rand"
)

var errEmptySlice = errors.New("string slice is empty")

func randomString(s []string) (string, error) {
	if len(s) < 1 {
		return "", errEmptySlice
	}

	r := rand.Intn(len(s))
	return s[r], nil
}
