package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileItemTemplate = "testdata/itemtemplate_test.json"

func TestItemTemplateRandomAlias(t *testing.T) {
	rand.Seed(1)

	it, err := loadItemTemplates(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := it["figurine"]
	if !ok {
		t.Fatal("cannot find 'figurine' item template")
	}

	a, err := fig.randomAlias()
	if err != nil {
		t.Fatal(err)
	}

	if a != "effigy" {
		t.Errorf("got: <%v>, want: <%v>", a, "effigy")
	}
}

func TestItemTemplateRandomAliasEmpty(t *testing.T) {
	it := itemTemplate{
		Aliases: []string{},
	}

	_, err := it.randomAlias()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestItemTemplateRandomMaterial(t *testing.T) {
	rand.Seed(1)

	res, err := LoadResources(testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, testFileModifier)
	if err != nil {
		t.Fatal(err)
	}

	it, err := loadItemTemplates(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := it["figurine"]
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

func TestItemTemplateRandomMaterialEmpty(t *testing.T) {
	res := &Resources{}
	it := itemTemplate{
		MaterialVariants: []string{},
	}

	_, err := it.randomMaterial(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}

	res = &Resources{
		materials: map[string]material{},
	}
	it = itemTemplate{
		MaterialVariants: []string{"test"},
	}

	_, err = it.randomMaterial(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestItemTemplateRandomDetail(t *testing.T) {
	rand.Seed(1)

	res, err := LoadResources(testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, testFileModifier)
	if err != nil {
		t.Fatal(err)
	}

	it, err := loadItemTemplates(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}

	fig, ok := it["figurine"]
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

func TestItemTemplateRandomDetailEmpty(t *testing.T) {
	res := &Resources{}
	it := itemTemplate{
		DetailVariants: []string{},
	}

	_, err := it.randomDetail(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}

	res = &Resources{
		details: map[string]detail{},
	}
	it = itemTemplate{
		DetailVariants: []string{"test"},
	}

	_, err = it.randomDetail(res)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestReadItemTemplates(t *testing.T) {
	f, err := os.Open(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	it, err := readItemTemplates(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(it) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(it), 5)
	}

	if it["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", it["figurine"].Name, "figurine")
	}

	if len(it["doll"].MaterialVariants) != 2 {
		t.Errorf("got: <%v>, want: <%v>", len(it["doll"].MaterialVariants), 2)
	}
}

func TestReadItemTemplatesEmpty(t *testing.T) {
	var f *os.File

	_, err := readItemTemplates(f)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadItemTemplates(t *testing.T) {
	it, err := loadItemTemplates(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}

	if len(it) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(it), 5)
	}

	if it["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", it["figurine"].Name, "figurine")
	}

	if len(it["doll"].MaterialVariants) != 2 {
		t.Errorf("got: <%v>, want: <%v>", len(it["doll"].MaterialVariants), 2)
	}
}

func TestLoadItemTemplatesEmpty(t *testing.T) {
	_, err := loadItemTemplates("fake file name")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}
