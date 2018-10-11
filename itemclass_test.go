package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileItemClass = "testdata/itemclass_test.json"

func TestItemClassRandomVerb(t *testing.T) {
	rand.Seed(1)

	ic, err := loadItemClasses(testFileItemClass)
	if err != nil {
		t.Fatal(err)
	}

	v, err := ic["art"].randomVerb()
	if err != nil {
		t.Fatal(err)
	}

	if v != "made to resemble" {
		t.Errorf("got: <%v>, want: <%v>", v, "made to resemble")
	}
}

func TestItemClassRandomVerbEmpty(t *testing.T) {
	ic := itemClass{VerbVariants: []string{}}

	_, err := ic.randomVerb()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestReadItemClasses(t *testing.T) {
	f, err := os.Open(testFileItemClass)
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
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadItemClasses(t *testing.T) {
	ic, err := loadItemClasses(testFileItemClass)
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
		t.Error("got: <nil>, want: <error>")
	}
}
