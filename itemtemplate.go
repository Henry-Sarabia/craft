package craft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type itemTemplate struct {
	Name             string            `json:"name"`
	ItemClass        string            `json:"item_class"`
	BaseValue        float64           `json:"base_value"`  // Measured in gold pieces
	BaseWeight       float64           `json:"base_weight"` // Measured in pounds
	Aliases          []string          `json:"aliases"`
	MaterialVariants []string          `json:"material_variants"`
	DetailVariants   []detailReference `json:"detail_variants"`
}

func (it *itemTemplate) craftPrototype(res *Resources) (*itemPrototype, error) {
	name, err := it.getName()
	if err != nil {
		return nil, err
	}

	cl, ok := res.itemClasses[it.ItemClass]
	if !ok {
		return nil, errors.Errorf("cannot find item class '%s' in available resources", it.ItemClass)
	}

	mat, err := it.randomMaterial(res)
	if err != nil {
		return nil, err
	}

	d, err := it.randomDetails()
	if err != nil {
		return nil, err
	}

	dets, err := res.getDetails(d)
	if err != nil {
		return nil, err
	}

	return &itemPrototype{
		name:     name,
		value:    it.BaseValue,
		weight:   it.BaseWeight,
		class:    &cl,
		material: mat,
		details:  dets,
	}, nil
}

func (it *itemTemplate) getName() (string, error) {
	a, err := randomString(it.Aliases)
	if err != nil {
		return "", errors.Wrap(err, "itemTemplate aliases slice is empty")
	}

	return a, nil
}

func (it *itemTemplate) getQuality(res *Resources) (string, error) {
	mat, err := it.randomMaterial(res)
	if err != nil {
		return "", err
	}

	q, err := mat.randomQuality(res)
	if err != nil {
		return "", err
	}

	return q, nil
}

func (it *itemTemplate) getMaterial(res *Resources) (string, error) {
	mat, err := it.randomMaterial(res)
	if err != nil {
		return "", err
	}

	vrnt, err := mat.randomVariant()
	if err != nil {
		return "", err
	}

	return vrnt, nil
}

func (it *itemTemplate) getFormat(res *Resources) (string, error) {
	cl, ok := res.itemClasses[it.ItemClass]
	if !ok {
		return "", errors.Errorf("cannot find item class '%s' in available resources", it.ItemClass)
	}

	return cl.Format, nil
}

func (it *itemTemplate) getValue(res *Resources) (float64, error) {
	return 0, nil
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

func (it *itemTemplate) randomDetails() (map[string]string, error) {
	m := make(map[string]string)

	for _, d := range it.DetailVariants {
		r, err := d.random()
		if err != nil {
			return nil, errors.Wrap(err, "cannot get random detailReference variant from itemTemplate")
		}
		m[d.Label] = r
	}

	return m, nil
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
