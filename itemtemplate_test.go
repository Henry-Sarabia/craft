package craft

import (
	"os"
	"testing"
)

const testFileItemTemplate = "testdata/itemtemplate_test.json"

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
		t.Errorf("got: <%v>, want: <nil>", err)
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
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}
