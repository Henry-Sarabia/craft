package craft

type material struct {
	Name         string   `json:"name"`
	ValueFactor  float64  `json:"value_factor"`
	WeightFactor float64  `json:"weight_factor"`
	Variants     []string `json:"variants"`
}
