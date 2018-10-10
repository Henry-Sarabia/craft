package craft

import (
	"io"
	"os"
	"testing"
)

func TestReadResources(t *testing.T) {
	temp, err := os.Open(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}
	defer temp.Close()

	class, err := os.Open(testFileItemClass)
	if err != nil {
		t.Fatal(err)
	}
	defer class.Close()

	mat, err := os.Open(testFileMaterial)
	if err != nil {
		t.Fatal(err)
	}
	defer mat.Close()

	det, err := os.Open(testFileDetail)
	if err != nil {
		t.Fatal(err)
	}
	defer det.Close()

	mod, err := os.Open(testFileModifier)
	if err != nil {
		t.Fatal(err)
	}
	defer mod.Close()

	res, err := ReadResources(temp, class, mat, det, mod)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.itemTemplates) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(res.itemTemplates), 5)
	}

	if res.itemTemplates["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", res.itemTemplates["figurine"], "figurine")
	}

	if len(res.itemClasses) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(res.itemClasses), 3)
	}

	if res.itemClasses["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", res.itemClasses["art"].Name, "art")
	}

	if len(res.materials) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(res.materials), 8)
	}

	if res.materials["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", res.materials["wood"].Name, "wood")
	}

	if len(res.details) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(res.details), 5)
	}

	if res.details["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", res.details["beverage"].Name, "beverage")
	}

	if len(res.modifiers) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(res.modifiers), 7)
	}

	if res.modifiers["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", res.modifiers["wear"].Name, "wear")
	}

}

func TestReadResourcesEmpty(t *testing.T) {
	var emptyFile *os.File

	temp, err := os.Open(testFileItemTemplate)
	if err != nil {
		t.Fatal(err)
	}
	defer temp.Close()

	class, err := os.Open(testFileItemClass)
	if err != nil {
		t.Fatal(err)
	}
	defer class.Close()

	mat, err := os.Open(testFileMaterial)
	if err != nil {
		t.Fatal(err)
	}
	defer mat.Close()

	det, err := os.Open(testFileDetail)
	if err != nil {
		t.Fatal(err)
	}
	defer det.Close()

	mod, err := os.Open(testFileModifier)
	if err != nil {
		t.Fatal(err)
	}
	defer mod.Close()

	var tests = []struct {
		Name  string
		Temp  io.Reader
		Class io.Reader
		Mat   io.Reader
		Det   io.Reader
		Mod   io.Reader
	}{
		{"Empty itemTemplate reader", emptyFile, class, mat, det, mod},
		{"Empty itemClass reader", temp, emptyFile, mat, det, mod},
		{"Empty material reader", temp, class, emptyFile, det, mod},
		{"Empty detail reader", temp, class, mat, emptyFile, mod},
		{"Empty modifier reader", temp, class, mat, det, emptyFile},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := ReadResources(tt.Temp, tt.Class, tt.Mat, tt.Det, tt.Mod)
			if err == nil {
				t.Errorf("got: <%v>, want: <nil>", err)
			}
		})
	}
}

func TestLoadResources(t *testing.T) {
	res, err := LoadResources(testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, testFileModifier)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.itemTemplates) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(res.itemTemplates), 5)
	}

	if res.itemTemplates["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", res.itemTemplates["figurine"], "figurine")
	}

	if len(res.itemClasses) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(res.itemClasses), 3)
	}

	if res.itemClasses["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", res.itemClasses["art"].Name, "art")
	}

	if len(res.materials) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(res.materials), 8)
	}

	if res.materials["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", res.materials["wood"].Name, "wood")
	}

	if len(res.details) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(res.details), 5)
	}

	if res.details["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", res.details["beverage"].Name, "beverage")
	}

	if len(res.modifiers) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(res.modifiers), 7)
	}

	if res.modifiers["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", res.modifiers["wear"].Name, "wear")
	}
}

func TestLoadResourcesEmpty(t *testing.T) {
	fakeFile := "fake file name"
	var tests = []struct {
		Name  string
		Temp  string
		Class string
		Mat   string
		Det   string
		Mod   string
	}{
		{"Non-existant itemTemplate file", fakeFile, testFileItemClass, testFileMaterial, testFileDetail, testFileModifier},
		{"Non-existant itemClass file", testFileItemTemplate, fakeFile, testFileMaterial, testFileDetail, testFileModifier},
		{"Non-existant material file", testFileItemTemplate, testFileItemClass, fakeFile, testFileDetail, testFileModifier},
		{"Non-existant detail file", testFileItemTemplate, testFileItemClass, testFileMaterial, fakeFile, testFileModifier},
		{"Non-existant modifier file", testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, fakeFile},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := LoadResources(tt.Temp, tt.Class, tt.Mat, tt.Det, tt.Mod)
			if err == nil {
				t.Errorf("got: <%v>, want: <nil>", err)
			}
		})
	}
}
