package craft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type template struct {
	Name             string            `json:"name"`
	Class            string            `json:"item_class"`
	BaseValue        float64           `json:"base_value"`  // Measured in gold pieces
	BaseWeight       float64           `json:"base_weight"` // Measured in pounds
	Aliases          []string          `json:"aliases"`
	MaterialVariants []string          `json:"material_variants"`
	DetailVariants   []detailReference `json:"detail_variants"`
}

func (t *template) randomName() (string, error) {
	a, err := randomString(t.Aliases)
	if err != nil {
		return "", errors.Wrap(err, "template aliases slice is empty")
	}

	return a, nil
}

func (t *template) randomMaterial() (string, error) {
	mat, err := randomString(t.MaterialVariants)
	if err != nil && err != errEmptySlice {
		return "", errors.Wrap(err, "cannot get random material from template")
	}

	return mat, nil
}

func (t *template) randomDetails() (map[string]string, error) {
	m := make(map[string]string)

	for _, d := range t.DetailVariants {
		r, err := d.randomVariant()
		if err != nil && err != errEmptySlice {
			return nil, errors.Wrap(err, "cannot get random detailReference variant from template")
		}
		m[d.Label] = r
	}

	return m, nil
}

func readTemplates(r io.Reader) (map[string]template, error) {
	var it []template

	if err := json.NewDecoder(r).Decode(&it); err != nil {
		return nil, err
	}

	m := make(map[string]template)
	for _, v := range it {
		m[v.Name] = v
	}

	return m, nil
}

func loadTemplates(filename string) (map[string]template, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readTemplates(f)
}
