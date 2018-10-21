package craft

import (
	"math/rand"
	"testing"
)

func TestDetailReferenceRandomVariant(t *testing.T) {
	rand.Seed(1)

	dr := detailReference{
		Label:    "shape",
		Variants: []string{"beast", "monstrosity"},
	}

	v, err := dr.randomVariant()
	if err != nil {
		t.Fatal(err)
	}

	if v != "monstrosity" {
		t.Errorf("got: <%v>, want: <%v>", v, "monstrosity")
	}
}

func TestDetailReferenceRandomVariantEmpty(t *testing.T) {
	dr := detailReference{
		Variants: []string{},
	}

	_, err := dr.randomVariant()
	if err == nil {
		t.Error("got: <nil>, want <error>")
	}
}
