package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileModifier = "testdata/modifier_test.json"

func TestModifierRandomVariant(t *testing.T) {
	rand.Seed(1)

	d, err := loadModifiers(testFileModifier)
	if err != nil {
		t.Fatal(err)
	}

	wear, err := d["wear"].randomVariant()
	if err != nil {
		t.Fatal(err)
	}

	if wear != "considerably disfigured" {
		t.Errorf("got: <%v>, want: <%v>", wear, "considerably disfigured")
	}
}

func TestModifierRandomVariantEmpty(t *testing.T) {
	m := modifier{Name: "empty", Variants: []string{}}

	_, err := m.randomVariant()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestReadModifiers(t *testing.T) {
	f, err := os.Open(testFileModifier)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	m, err := readModifiers(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(m) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(m), 7)
	}

	if m["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", m["wear"].Name, "wear")
	}

	if len(m["composite"].Variants) != 19 {
		t.Errorf("got: <%v>, want: <%v>", len(m["composite"].Variants), 19)
	}
}

func TestReadModifiersEmpty(t *testing.T) {
	var f *os.File

	_, err := readModifiers(f)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadModifiers(t *testing.T) {
	m, err := loadModifiers(testFileModifier)
	if err != nil {
		t.Fatal(err)
	}

	if len(m) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(m), 7)
	}

	if m["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", m["wear"].Name, "wear")
	}

	if len(m["composite"].Variants) != 19 {
		t.Errorf("got: <%v>, want: <%v>", len(m["composite"].Variants), 19)
	}
}

func TestLoadModifiersEmpty(t *testing.T) {
	_, err := loadModifiers("fake file name")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}
