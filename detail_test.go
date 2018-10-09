package craft

import (
	"os"
	"testing"
)

const testDetailFile = "testdata/detail_test.json"

func TestReadDetails(t *testing.T) {
	f, err := os.Open(testDetailFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d, err := readDetails(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(d) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(d), 5)
	}

	if d["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", d["beverage"].Name, "beverage")
	}

	if len(d["humanoid"].Variants) != 23 {
		t.Errorf("got: <%v>, want: <%v>", len(d["humanoid"].Variants), 23)
	}
}

func TestReadDetailsEmpty(t *testing.T) {
	var f *os.File

	_, err := readDetails(f)
	if err == nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}

func TestLoadDetails(t *testing.T) {
	d, err := loadDetails(testDetailFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(d) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(d), 5)
	}

	if d["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", d["beverage"].Name, "beverage")
	}

	if len(d["humanoid"].Variants) != 23 {
		t.Errorf("got: <%v>, want: <%v>", len(d["humanoid"].Variants), 23)
	}
}

func TestLoadDetailsError(t *testing.T) {
	_, err := loadDetails("fake file name")
	if err == nil {
		t.Errorf("got: <%v>, want: <nil>", err)
	}
}
