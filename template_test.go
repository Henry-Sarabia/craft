package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileTemplate = "testdata/template_test.json"

func TestTemplateRandomName(t *testing.T) {
	rand.Seed(1)

	tmp, err := loadTemplates(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := tmp["figurine"]
	if !ok {
		t.Fatal("cannot find 'figurine' item template")
	}

	n, err := fig.randomName()
	if err != nil {
		t.Fatal(err)
	}

	if n != "figure" {
		t.Errorf("got: <%v>, want: <%v>", n, "figure")
	}
}

func TestTemplateRandomNameEmpty(t *testing.T) {
	tmp := template{
		Aliases: []string{},
	}

	_, err := tmp.randomName()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestTemplateRandomMaterial(t *testing.T) {
	rand.Seed(1)

	tmp, err := loadTemplates(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := tmp["figurine"]
	if !ok {
		t.Fatal("cannot find 'figurine' item template")
	}

	m, err := fig.randomMaterial()
	if err != nil {
		t.Fatal(err)
	}

	if m != "wood" {
		t.Errorf("got: <%v>, want: <%v>", m, "wood")
	}
}

func TestTemplateRandomMaterialEmpty(t *testing.T) {
	tmp := template{
		MaterialVariants: []string{},
	}

	_, err := tmp.randomMaterial()
	if err != nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}

func TestTemplateRandomDetails(t *testing.T) {
	rand.Seed(1)

	tmp, err := loadTemplates(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := tmp["figurine"]
	if !ok {
		t.Fatal("cannot find 'figurine' item template")
	}

	d, err := fig.randomDetails()
	if err != nil {
		t.Fatal(err)
	}

	if d["shape"] != "monstrosity" {
		t.Errorf("got: <%v>, want: <%v>", d["shape"], "monstrosity")
	}
}

func TestTemplateRandomDetailEmpty(t *testing.T) {
	tmp := template{
		DetailVariants: []detailReference{},
	}

	_, err := tmp.randomDetails()
	if err != nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}

func TestReadTemplates(t *testing.T) {
	f, err := os.Open(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	tmp, err := readTemplates(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(tmp) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(tmp), 5)
	}

	if tmp["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", tmp["figurine"].Name, "figurine")
	}

	if len(tmp["doll"].MaterialVariants) != 2 {
		t.Errorf("got: <%v>, want: <%v>", len(tmp["doll"].MaterialVariants), 2)
	}
}

func TestReadTemplatesEmpty(t *testing.T) {
	var f *os.File

	_, err := readTemplates(f)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadTemplates(t *testing.T) {
	tmp, err := loadTemplates(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}

	if len(tmp) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(tmp), 5)
	}

	if tmp["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", tmp["figurine"].Name, "figurine")
	}

	if len(tmp["doll"].MaterialVariants) != 2 {
		t.Errorf("got: <%v>, want: <%v>", len(tmp["doll"].MaterialVariants), 2)
	}
}

func TestLoadTemplatesEmpty(t *testing.T) {
	_, err := loadTemplates("fake file name")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}
