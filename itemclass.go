package craft

import (
	"encoding/json"
	"io"
	"os"
)

type itemClass struct {
	Name         string   `json:"name"`
	Format       string   `json:"format"`
	Example      string   `json:"example"`
	VerbVariants []string `json:"verb_variants"`
}

func readItemClasses(r io.Reader) (map[string]itemClass, error) {
	var ic []itemClass

	if err := json.NewDecoder(r).Decode(&ic); err != nil {
		return nil, err
	}

	m := make(map[string]itemClass)
	for _, v := range ic {
		m[v.Name] = v
	}

	return m, nil
}

func loadItemClasses(filename string) (map[string]itemClass, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readItemClasses(f)
}
