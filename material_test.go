package craft

import (
	"os"
	"testing"
)

const testMaterialFile = "testdata/material_test.json"

func TestReadMaterials(t *testing.T) {
	f, err := os.Open(testMaterialFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	m, err := readMaterials(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(m) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(m), 8)
	}

	if m["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", m["wood"].Name, "wood")
	}

	if len(m["animal skin"].Variants) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(m["animal skin"].Variants), 7)
	}
}

func TestReadMaterialsEmpty(t *testing.T) {
	var f *os.File

	_, err := readMaterials(f)
	if err == nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}

func TestLoadMaterials(t *testing.T) {
	m, err := loadMaterials(testMaterialFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(m) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(m), 8)
	}

	if m["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", m["wood"].Name, "wood")
	}

	if len(m["animal skin"].Variants) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(m["animal skin"].Variants), 7)
	}
}

func TestLoadMaterialsEmpty(t *testing.T) {
	_, err := loadMaterials("fake file name")
	if err == nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}