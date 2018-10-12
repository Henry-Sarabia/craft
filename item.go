package craft

import (
	"strings"

	"github.com/Henry-Sarabia/article"
	"github.com/pkg/errors"
)

// Item represents an item that would commonly be found in a medieval RPG game.
type Item struct {
	Name     string  `json:"name"`
	Format   string  `json:"format"`
	Quality  string  `json:"quality"`
	Material string  `json:"material"`
	Detail   string  `json:"detail"`
	Verb     string  `json:"verb"`
	Value    float64 `json:"value"`
	Weight   float64 `json:"weight"`

	Description string `json:"description"`
}

// generateItem generates a new item using the provided template and resources.
func generateItem(tmp *itemTemplate, res *Resources) (*Item, error) {
	proto, err := tmp.craftPrototype(res)
	if err != nil {
		return nil, err
	}

	i, err := proto.craftItem(res)
	if err != nil {
		return nil, err
	}

	err = i.composeDescription()
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Item) composeDescription() error {
	tok, err := tokenize(i.Format)
	if err != nil {
		return err
	}

	i.Description, err = i.parse(tok)
	if err != nil {
		return err
	}

	return nil
}

func tokenize(format string) ([]string, error) {
	tok := strings.Fields(format)
	if tok[len(tok)-1] == "<article>" {
		return nil, errors.New("article token cannot be the last token in a format")
	}
	return tok, nil
}

func (i *Item) parse(tok []string) (string, error) {
	for j := len(tok) - 1; j >= 0; j-- {
		switch tok[j] {
		case "<article>":
			tok[j] = article.Indefinite(tok[j+1])

		case "<quality>":
			tok[j] = i.Quality

		case "<material>":
			tok[j] = i.Material

		case "<name>":
			tok[j] = i.Name

		case "<verb>":
			tok[j] = i.Verb

		case "<detail>":
			tok[j] = i.Detail

		default:
			return "", errors.Errorf("unexpected token '%v' in item format", tok[j])
		}
	}

	return strings.Join(tok, " "), nil
}

// func parse(tok []string, tokMap map[string]string) (string, error) {
// 	for i := len(tok) - 1; i >= 0; i-- {
// 		cur, ok := tokMap[tok[i]]

// 		if tok[i]
// 	}
// }
