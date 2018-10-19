package craft

import "github.com/pkg/errors"

type detailReference struct {
	Label    string   `json:"label"`
	Variants []string `json:"variants"`
}

func (df *detailReference) random() (string, error) {
	v, err := randomString(df.Variants)
	if err != nil {
		return "", errors.Wrap(err, "detailReference variants slice is empty")
	}

	return v, nil
}
