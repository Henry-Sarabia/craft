package craft

import (
	"encoding/json"
	"io"
	"os"
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
