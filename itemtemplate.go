package craft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type itemTemplate struct {
	Name             string   `json:"name"`
	ItemClass        string   `json:"item_class"`
	BaseValue        float64  `json:"base_value"`  // Measured in gold pieces
	BaseWeight       float64  `json:"base_weight"` // Measured in pounds
	Aliases          []string `json:"aliases"`
	MaterialVariants []string `json:"material_variants"`
	DetailVariants   []string `json:"detail_variants"`
}

// func NewItem(temp itemTemplate, pool *ResourcePool) (*Item, error) {
// 	return nil, nil
// }

func (it *itemTemplate) randomAlias() (string, error) {
	a, err := randomString(it.Aliases)
	if err != nil {
		return "", errors.Wrap(err, "itemTemplate aliases slice is empty")
	}

	return a, nil
}

func (it *itemTemplate) randomMaterial(res *Resources) (*material, error) {
	matName, err := randomString(it.MaterialVariants)
	if err != nil {
		return nil, errors.Wrap(err, "itemTemplate material variants slice is empty")
	}

	mat, ok := res.materials[matName]
	if !ok {
		return nil, errors.Errorf("cannot find material '%s' in available resources", matName)
	}

	return &mat, nil
}

func (it *itemTemplate) randomDetail(res *Resources) (*detail, error) {
	detName, err := randomString(it.DetailVariants)
	if err != nil {
		return nil, errors.Wrap(err, "itemTemplate detail variants slice is empty")
	}

	det, ok := res.details[detName]
	if !ok {
		return nil, errors.Errorf("cannot find detail '%s' in available resources", detName)
	}

	return &det, nil
}

func readItemTemplates(r io.Reader) (map[string]itemTemplate, error) {
	var it []itemTemplate

	if err := json.NewDecoder(r).Decode(&it); err != nil {
		return nil, err
	}

	m := make(map[string]itemTemplate)
	for _, v := range it {
		m[v.Name] = v
	}

	return m, nil
}

func loadItemTemplates(filename string) (map[string]itemTemplate, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readItemTemplates(f)
}
