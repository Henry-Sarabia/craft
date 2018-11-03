package craft

import "math"

type prototype struct {
	name     string
	value    float64
	weight   float64
	class    string
	config   *configuration
	material *material
	details  map[string]*detail
}

func (p *prototype) getMaterial() (string, error) {
	m, err := p.material.randomVariant()
	if err != nil {
		return "", err
	}

	return m, nil
}

func (p *prototype) getQuality(m map[string]quality) (string, error) {
	q, err := p.material.randomQuality(m)
	if err != nil {
		return "", err
	}

	return q, nil
}

func (p *prototype) getDetails() (map[string]string, error) {
	m := make(map[string]string)

	for lbl, d := range p.details {
		det, err := d.randomVariant()
		if err != nil {
			return nil, err
		}

		m[lbl] = det
	}

	return m, nil
}

func (p *prototype) getFormat() (string, error) {
	return p.config.Format, nil
}

func (p *prototype) getVerb() (string, error) {
	return p.config.randomVerb()
}

func (p *prototype) getValue() float64 {
	v := p.value * p.material.ValueFactor
	return preciseRound(v, 1)
}

func (p *prototype) getWeight() float64 {
	w := p.weight * p.material.WeightFactor
	return preciseRound(w, 1)
}

func preciseRound(num float64, pre int) float64 {
	x := math.Pow(10, float64(pre))
	return float64(math.Round(num*x)) / x
}
