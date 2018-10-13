package craft

import (
	"math/rand"
	"os"
	"testing"
)

const testFileQuality = "testdata/quality_test.json"

func TestQualityRandomVariant(t *testing.T) {
	rand.Seed(1)

	d, err := loadQualities(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}

	wear, err := d["wear"].randomVariant()
	if err != nil {
		t.Fatal(err)
	}

	if wear != "considerably disfigured" {
		t.Errorf("got: <%v>, want: <%v>", wear, "considerably disfigured")
	}
}

func TestQualityRandomVariantEmpty(t *testing.T) {
	q := quality{Name: "empty", Variants: []string{}}

	_, err := q.randomVariant()
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestReadQualities(t *testing.T) {
	f, err := os.Open(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q, err := readQualities(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(q) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(q), 7)
	}

	if q["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", q["wear"].Name, "wear")
	}

	if len(q["composite"].Variants) != 19 {
		t.Errorf("got: <%v>, want: <%v>", len(q["composite"].Variants), 19)
	}
}

func TestReadQualitiesEmpty(t *testing.T) {
	var f *os.File

	_, err := readQualities(f)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestLoadQualities(t *testing.T) {
	q, err := loadQualities(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}

	if len(q) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(q), 7)
	}

	if q["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", q["wear"].Name, "wear")
	}

	if len(q["composite"].Variants) != 19 {
		t.Errorf("got: <%v>, want: <%v>", len(q["composite"].Variants), 19)
	}
}

func TestLoadQualitiesEmpty(t *testing.T) {
	_, err := loadQualities("fake file name")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}
