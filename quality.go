package craft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type quality struct {
	Name     string   `json:"name"`
	Variants []string `json:"variants"`
}

func (q quality) randomVariant() (string, error) {
	v, err := randomString(q.Variants)
	if err != nil {
		return "", errors.Wrap(err, "quality variants slice is empty")
	}

	return v, nil
}

func readQualities(r io.Reader) (map[string]quality, error) {
	var qual []quality

	if err := json.NewDecoder(r).Decode(&qual); err != nil {
		return nil, err
	}

	m := make(map[string]quality)
	for _, v := range qual {
		m[v.Name] = v
	}

	return m, nil
}

func loadQualities(filename string) (map[string]quality, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readQualities(f)
}
