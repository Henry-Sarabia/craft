package craft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type class struct {
	Name         string   `json:"name"`
	Format       string   `json:"format"`
	Example      string   `json:"example"`
	VerbVariants []string `json:"verb_variants"`
}

func (cl class) randomVerb() (string, error) {
	v, err := randomString(cl.VerbVariants)
	if err != nil {
		return "", errors.Wrap(err, "class VerbVariants slice is empty")
	}

	return v, nil
}

func readClasses(r io.Reader) (map[string]class, error) {
	var cl []class

	if err := json.NewDecoder(r).Decode(&cl); err != nil {
		return nil, err
	}

	m := make(map[string]class)
	for _, v := range cl {
		m[v.Name] = v
	}

	return m, nil
}

func loadClasses(filename string) (map[string]class, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readClasses(f)
}
