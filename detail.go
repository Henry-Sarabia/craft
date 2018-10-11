package craft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type detail struct {
	Name     string   `json:"name"`
	Variants []string `json:"variants"`
}

func (d detail) randomVariant() (string, error) {
	v, err := randomString(d.Variants)
	if err != nil {
		return "", errors.Wrap(err, "detail variants slice is empty")
	}

	return v, nil
}

func readDetails(r io.Reader) (map[string]detail, error) {
	var d []detail

	if err := json.NewDecoder(r).Decode(&d); err != nil {
		return nil, err
	}

	m := make(map[string]detail)
	for _, v := range d {
		m[v.Name] = v
	}

	return m, nil
}

func loadDetails(filename string) (map[string]detail, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readDetails(f)
}
