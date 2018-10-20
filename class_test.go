package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileClass = "testdata/class_test.json"

func TestClassRandomVerb(t *testing.T) {
	rand.Seed(1)

	cl, err := loadClasses(testFileClass)
	if err != nil {
		t.Fatal(err)
	}

	v, err := cl["art"].randomVerb()
	if err != nil {
		t.Fatal(err)
	}

	if v != "made to resemble" {
		t.Errorf("got: <%v>, want: <%v>", v, "made to resemble")
	}
}

func TestClassRandomVerbEmpty(t *testing.T) {
	cl := class{VerbVariants: []string{}}

	_, err := cl.randomVerb()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestReadClasses(t *testing.T) {
	f, err := os.Open(testFileClass)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	cl, err := readClasses(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(cl) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(cl), 3)
	}

	if cl["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", cl["art"].Name, "art")
	}

	if len(cl["jewelry"].VerbVariants) != 6 {
		t.Errorf("got: <%v>, want: <%v>", len(cl["jewelry"].VerbVariants), 6)
	}
}

func TestReadClassesEmpty(t *testing.T) {
	var f *os.File

	_, err := readClasses(f)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadClasses(t *testing.T) {
	cl, err := loadClasses(testFileClass)
	if err != nil {
		t.Fatal(err)
	}

	if len(cl) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(cl), 3)
	}

	if cl["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", cl["art"].Name, "art")
	}

	if len(cl["jewelry"].VerbVariants) != 6 {
		t.Errorf("got: <%v>, want: <%v>", len(cl["jewelry"].VerbVariants), 6)
	}
}

func TestLoadClassesEmpty(t *testing.T) {
	_, err := loadClasses("fake file name")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}
