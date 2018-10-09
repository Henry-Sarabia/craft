package craft

import (
	"os"
	"testing"
)

const testModifierFile = "testdata/modifier_test.json"

func TestReadModifiers(t *testing.T) {
	f, err := os.Open(testModifierFile)
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
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}

func TestLoadModifiers(t *testing.T) {
	m, err := loadModifiers(testModifierFile)
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
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}
