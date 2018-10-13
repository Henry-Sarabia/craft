package craft

import (
	"os"
	"testing"
)

func TestResourcesSelectTemplate(t *testing.T) {
	r, err := LoadResources(testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, testFileQuality)
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := r.selectTemplate()
	if err != nil {
		t.Fatal(err)
	}

	if tmp.Name == "" {
		t.Errorf("got: <%v>, want: <%v>", "", tmp.Name)
	}
}

func TestResourcesSelectTemplateEmpty(t *testing.T) {
	r := Resources{
		itemTemplates: map[string]itemTemplate{},
	}

	_, err := r.selectTemplate()
	if err != errEmptyTemplateMap {
		t.Errorf("got: <%v>, want: <%v>", nil, errEmptyTemplateMap)
	}
}

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

	qual, err := os.Open(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}
	defer qual.Close()

	res, err := ReadResources(temp, class, mat, det, qual)
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

	if len(res.qualities) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(res.qualities), 7)
	}

	if res.qualities["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", res.qualities["wear"].Name, "wear")
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

	qual, err := os.Open(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}
	defer qual.Close()

	var tests = []struct {
		Name  string
		Temp  *os.File
		Class *os.File
		Mat   *os.File
		Det   *os.File
		Qual  *os.File
	}{
		{"Empty itemTemplate reader", emptyFile, class, mat, det, qual},
		{"Empty itemClass reader", temp, emptyFile, mat, det, qual},
		{"Empty material reader", temp, class, emptyFile, det, qual},
		{"Empty detail reader", temp, class, mat, emptyFile, qual},
		{"Empty quality reader", temp, class, mat, det, emptyFile},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := ReadResources(tt.Temp, tt.Class, tt.Mat, tt.Det, tt.Qual)
			if err == nil {
				t.Error("got: <nil>, want: <error>")
			}

			// Rewind to beginning of files before next run
			tt.Temp.Seek(0, 0)
			tt.Class.Seek(0, 0)
			tt.Mat.Seek(0, 0)
			tt.Det.Seek(0, 0)
			tt.Qual.Seek(0, 0)
		})
	}
}

func TestLoadResources(t *testing.T) {
	res, err := LoadResources(testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, testFileQuality)
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

	if len(res.qualities) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(res.qualities), 7)
	}

	if res.qualities["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", res.qualities["wear"].Name, "wear")
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
		Qual  string
	}{
		{"Non-existant itemTemplate file", fakeFile, testFileItemClass, testFileMaterial, testFileDetail, testFileQuality},
		{"Non-existant itemClass file", testFileItemTemplate, fakeFile, testFileMaterial, testFileDetail, testFileQuality},
		{"Non-existant material file", testFileItemTemplate, testFileItemClass, fakeFile, testFileDetail, testFileQuality},
		{"Non-existant detail file", testFileItemTemplate, testFileItemClass, testFileMaterial, fakeFile, testFileQuality},
		{"Non-existant quality file", testFileItemTemplate, testFileItemClass, testFileMaterial, testFileDetail, fakeFile},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := LoadResources(tt.Temp, tt.Class, tt.Mat, tt.Det, tt.Qual)
			if err == nil {
				t.Error("got: <nil>, want: <error>")
			}
		})
	}
}
