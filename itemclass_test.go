package craft

import (
	"os"
	"testing"
)

const testItemClassFile = "testdata/itemclass_test.json"

func TestReadItemClasses(t *testing.T) {
	f, err := os.Open(testItemClassFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	ic, err := readItemClasses(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(ic) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(ic), 3)
	}

	if ic["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", ic["art"].Name, "art")
	}

	if len(ic["jewelry"].VerbVariants) != 6 {
		t.Errorf("got: <%v>, want: <%v>", len(ic["jewelry"].VerbVariants), 6)
	}
}

func TestReadItemClassesEmpty(t *testing.T) {
	var f *os.File

	_, err := readItemClasses(f)
	if err == nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}

func TestLoadItemClasses(t *testing.T) {
	ic, err := loadItemClasses(testItemClassFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(ic) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(ic), 3)
	}

	if ic["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", ic["art"].Name, "art")
	}

	if len(ic["jewelry"].VerbVariants) != 6 {
		t.Errorf("got: <%v>, want: <%v>", len(ic["jewelry"].VerbVariants), 6)
	}
}

func TestLoadItemClassesEmpty(t *testing.T) {
	_, err := loadItemClasses("fake file name")
	if err == nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}
