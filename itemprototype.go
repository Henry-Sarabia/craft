package craft

type itemPrototype struct {
	name     string
	value    float64
	weight   float64
	class    *itemClass
	material *material
	detail   *detail
}

func (ip *itemPrototype) craftItem(res *Resources) (*Item, error) {
	i := &Item{}

	var err error

	if i.Quality, err = ip.getQuality(res); err != nil {
		return nil, err
	}

	if i.Material, err = ip.getMaterial(); err != nil {
		return nil, err
	}

	if i.Detail, err = ip.getDetail(); err != nil {
		return nil, err
	}

	if i.Format, err = ip.getFormat(); err != nil {
		return nil, err
	}

	if i.Verb, err = ip.getVerb(); err != nil {
		return nil, err
	}

	i.Name = ip.name
	i.Value = ip.getValue()
	i.Weight = ip.getWeight()

	return i, nil
}

func (ip *itemPrototype) getQuality(res *Resources) (string, error) {
	q, err := ip.material.randomModifier(res)
	if err != nil {
		return "", err
	}

	return q, nil
}

func (ip *itemPrototype) getMaterial() (string, error) {
	m, err := ip.material.randomVariant()
	if err != nil {
		return "", err
	}

	return m, nil
}

func (ip *itemPrototype) getDetail() (string, error) {
	d, err := ip.detail.randomVariant()
	if err != nil {
		return "", err
	}

	return d, nil
}

func (ip *itemPrototype) getFormat() (string, error) {
	return ip.class.Format, nil
}

func (ip *itemPrototype) getVerb() (string, error) {
	return ip.class.randomVerb()
}

func (ip *itemPrototype) getValue() float64 {
	return ip.value * ip.material.ValueFactor
}

func (ip *itemPrototype) getWeight() float64 {
	return ip.weight * ip.material.WeightFactor
}
