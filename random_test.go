package craft

import (
	"math/rand"
	"testing"
)

func TestRandomString(t *testing.T) {
	rand.Seed(1)

	s := []string{"one", "two", "three"}

	r, err := randomString(s)
	if err != nil {
		t.Fatal(err)
	}

	if r != "three" {
		t.Errorf("got: <%v>, want: <%v>", r, "three")
	}
}

func TestRandomStringEmpty(t *testing.T) {
	empty := []string{}

	_, err := randomString(empty)
	if err != errEmptySlice {
		t.Errorf("got: <%v>, want: <%v>", nil, errEmptySlice)
	}
}
