package craft

import (
	"encoding/json"
	"io"
	"os"
)

type material struct {
	Name         string   `json:"name"`
	ValueFactor  float64  `json:"value_factor"`
	WeightFactor float64  `json:"weight_factor"`
	Modifiers    []string `json:"modifiers"`
	Variants     []string `json:"variants"`
}

func readMaterials(r io.Reader) (map[string]material, error) {
	var mat []material

	if err := json.NewDecoder(r).Decode(&mat); err != nil {
		return nil, err
	}

	m := make(map[string]material)
	for _, v := range mat {
		m[v.Name] = v
	}

	return m, nil
}

func loadMaterials(filename string) (map[string]material, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readMaterials(f)
}
