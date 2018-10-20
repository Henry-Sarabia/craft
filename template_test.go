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

	a, err := fig.randomName()
	if err != nil {
		t.Fatal(err)
	}

	if a != "effigy" {
		t.Errorf("got: <%v>, want: <%v>", a, "effigy")
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

//func TestTemplateRandomQuality(t *testing.T) {
//	rand.Seed(1)
//
//	tmp, err := loadTemplates(testFileTemplate)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fig, ok := tmp["figurine"]
//	if !ok {
//		t.Fatal("cannot find 'figurine' item template")
//	}
//
//	q, err := fig.randomQuality(c)
//	if err != nil {
//
//	}
//}

func TestTemplateRandomMaterial(t *testing.T) {
	rand.Seed(1)

	res, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := loadTemplates(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := tmp["figurine"]
	if !ok {
		t.Fatal("cannot find 'figurine' item template")
	}

	m, err := fig.randomMaterial(res)
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", m.Name, "wood")
	}
}

func TestTemplateRandomMaterialEmpty(t *testing.T) {
	res := &Crafter{}
	tmp := template{
		MaterialVariants: []string{},
	}

	_, err := tmp.randomMaterial(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}

	res = &Crafter{
		materials: map[string]material{},
	}
	tmp = template{
		MaterialVariants: []string{"test"},
	}

	_, err = tmp.randomMaterial(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestTemplateRandomDetail(t *testing.T) {
	rand.Seed(1)

	res, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := loadTemplates(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := tmp["figurine"]
	if !ok {
		t.Fatal("cannot find 'figurine' item template")
	}

	d, err := fig.randomDetail(res)
	if err != nil {
		t.Fatal(err)
	}

	if d.Name != "monstrosity" {
		t.Errorf("got: <%v>, want: <%v>", d.Name, "monstrosity")
	}
}

func TestTemplateRandomDetailEmpty(t *testing.T) {
	res := &Crafter{}
	tmp := template{
		DetailVariants: []string{},
	}

	_, err := tmp.randomDetail(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}

	res = &Crafter{
		details: map[string]detail{},
	}
	tmp = template{
		DetailVariants: []string{"test"},
	}

	_, err = tmp.randomDetail(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
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
