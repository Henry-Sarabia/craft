package craft

type itemTemplate struct {
	Name             string    `json:"name"`
	ItemClass        itemClass `json:"item_class"`
	BaseValue        float64   `json:"base_value"`  // Measured in gold pieces
	BaseWeight       float64   `json:"base_weight"` // Measured in pounds
	Aliases          []string  `json:"aliases"`
	MaterialVariants []string  `json:"material_variants"`
	DetailVariants   []string  `json:"detail_variants"`
}

func NewItem(temp itemTemplate, pool *ResourcePool) (*Item, error) {
	return nil, nil
}
