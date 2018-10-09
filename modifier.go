package craft

import (
	"encoding/json"
	"io"
	"os"
)

type modifier struct {
	Name     string   `json:"name"`
	Variants []string `json:"variants"`
}

func readModifiers(r io.Reader) (map[string]modifier, error) {
	var mod []modifier

	if err := json.NewDecoder(r).Decode(&mod); err != nil {
		return nil, err
	}

	m := make(map[string]modifier)
	for _, v := range mod {
		m[v.Name] = v
	}

	return m, nil
}

func loadModifiers(filename string) (map[string]modifier, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readModifiers(f)
}
