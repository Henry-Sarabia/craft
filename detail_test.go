package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileDetail = "testdata/detail_test.json"

func TestDetailRandomVariant(t *testing.T) {
	rand.Seed(1)

	d, err := loadDetails(testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	bev, err := d["beverage"].randomVariant()
	if err != nil {
		t.Fatal(err)
	}

	if bev != "mulled mead" {
		t.Errorf("got: <%v>, want: <%v>", bev, "mulled mead")
	}
}

func TestDetailRandomVariantEmpty(t *testing.T) {
	d := detail{Name: "empty", Variants: []string{}}

	_, err := d.randomVariant()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestReadDetails(t *testing.T) {
	f, err := os.Open(testFileDetail)
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
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadDetails(t *testing.T) {
	d, err := loadDetails(testFileDetail)
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
		t.Error("got: <nil>, want: <error>")
	}
}
