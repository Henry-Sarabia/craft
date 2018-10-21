package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileMaterial = "testdata/material_test.json"

func TestMaterialRandomVariant(t *testing.T) {
	rand.Seed(1)

	d, err := loadMaterials(testFileMaterial)
	if err != nil {
		t.Fatal(err)
	}

	wood, err := d["wood"].randomVariant()
	if err != nil {
		t.Fatal(err)
	}

	if wood != "corkwood" {
		t.Errorf("got: <%v>, want: <%v>", wood, "corkwood")
	}
}

func TestMaterialRandomVariantEmpty(t *testing.T) {
	m := material{Name: "empty", Variants: []string{}}

	_, err := m.randomVariant()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestMaterialRandomQuality(t *testing.T) {
	rand.Seed(1)

	d, err := loadMaterials(testFileMaterial)
	if err != nil {
		t.Fatal(err)
	}

	q, err := loadQualities(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}

	qual, err := d["wood"].randomQuality(q)
	if err != nil {
		t.Fatal(err)
	}

	if qual != "heavily charred" {
		t.Errorf("got: <%v>, want: <%v>", qual, "heavily charred")
	}
}

func TestMaterialRandomQualityEmpty(t *testing.T) {
	m := material{
		Qualities: []string{},
	}

	q := make(map[string]quality)

	_, err := m.randomQuality(q)
	if err == nil {
		t.Error("got: <nil>, want :<error>")
	}
}

func TestReadMaterials(t *testing.T) {
	f, err := os.Open(testFileMaterial)
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
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadMaterials(t *testing.T) {
	m, err := loadMaterials(testFileMaterial)
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
		t.Error("got: <nil>, want: <error>")
	}
}
