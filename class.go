package craft

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

type class struct {
	Name    string          `json:"name"`
	Configs []configuration `json:"configurations"`
}

type configuration struct {
	Name            string   `json:"name"`
	Format          string   `json:"format"`
	Example         string   `json:"example"`
	RequiredDetails []string `json:"required_details"`
	VerbVariants    []string `json:"verb_variants"`
}

func (cf configuration) randomVerb() (string, error) {
	v, err := randomString(cf.VerbVariants)
	if err != nil {
		return "", errors.Wrap(err, "configuration VerbVariants slice is empty")
	}

	return v, nil
}

func (cl class) randomConfiguration() (*configuration, error) {
	if len(cl.Configs) < 1 {
		return nil, errors.New("class configurations slice is empty")
	}

	r := rand.Intn(len(cl.Configs))
	return &cl.Configs[r], nil
}

func (cl class) getConfig(name string) (*configuration, error) {
	for _, cf := range cl.Configs {
		if cf.Name == name {
			return &cf, nil
		}
	}

	return nil, errors.Errorf("cannot find class configuration '%s'", name)
}

func readClasses(r io.Reader) (map[string]class, error) {
	var cl []class

	if err := json.NewDecoder(r).Decode(&cl); err != nil {
		return nil, err
	}

	m := make(map[string]class)
	for _, v := range cl {
		m[v.Name] = v
	}

	return m, nil
}

func loadClasses(filename string) (map[string]class, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readClasses(f)
}
