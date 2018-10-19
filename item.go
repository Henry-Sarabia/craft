package craft

import (
	"strings"

	"github.com/Henry-Sarabia/article"
	"github.com/pkg/errors"
)

// Item represents an item that would commonly be found in a medieval RPG game.
type Item struct {
	Name     string            `json:"name"`
	Material string            `json:"material"`
	Quality  string            `json:"quality"`
	Details  map[string]string `json:"detail"`
	Format   string            `json:"format"`
	Verb     string            `json:"verb"`
	Value    float64           `json:"value"`
	Weight   float64           `json:"weight"`

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

	i.Description, err = i.replaceTokens(tok)
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

func (i *Item) replaceDetailTokens(tok []string) []string {
	for j, t := range tok {
		t = strings.Trim(t, "<>")
		if d, ok := i.Details[t]; ok {
			tok[j] = d
		}
	}

	return tok
}

func (i *Item) replaceTokens(tok []string) (string, error) {
	tok = i.replaceDetailTokens(tok)
	for j := len(tok) - 1; j >= 0; j-- {
		switch tok[j] {
		case "<name>":
			tok[j] = i.Name

		case "<material>":
			tok[j] = i.Material

		case "<quality>":
			tok[j] = i.Quality

		case "<article>":
			tok[j] = article.Indefinite(tok[j+1])

		case "<verb>":
			tok[j] = i.Verb
		}
	}

	return strings.Join(tok, " "), nil
}
