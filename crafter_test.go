package craft

import (
	"math/rand"
	"os"
	"testing"
)

func TestCrafterNewItem(t *testing.T) {
	rand.Seed(1)

	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	i, err := c.NewItem()
	if err != nil {
		t.Fatal(err)
	}

	if i.Name != "bota bag" {
		t.Errorf("got: <%v>, want: <%v>", i.Name, "bota bag")
	}

	if i.Value != 1.5 {
		t.Errorf("got: <%v>, want: <%v>", i.Value, 1.5)
	}

	if i.Weight != 1 {
		t.Errorf("got: <%v>, want: <%v>", i.Weight, 1)
	}
}

func TestCrafterGenerateItem(t *testing.T) {
	rand.Seed(1)

	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	fig := c.templates["figurine"]

	i, err := c.generateItem(&fig)
	if err != nil {
		t.Fatal(err)
	}

	if i.Value != 30 {
		t.Errorf("got: <%v>, want: <%v>", i.Value, 30)
	}

	if i.Weight != 9 {
		t.Errorf("got: <%v>, want: <%v>", i.Weight, 9)
	}
}

func TestCrafterToPrototype(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	fig := c.templates["figurine"]

	p, err := c.toPrototype(&fig)
	if err != nil {
		t.Fatal(err)
	}

	if p.value != 10 {
		t.Errorf("got: <%v>, want: <%v>", p.value, 10)
	}

	if p.weight != 1 {
		t.Errorf("got: <%v>, want: <%v>", p.weight, 1)
	}
}

func TestCrafterToItem(t *testing.T) {
	rand.Seed(1)

	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	fig := c.templates["figurine"]

	p, err := c.toPrototype(&fig)
	if err != nil {
		t.Fatal(err)
	}

	i, err := c.toItem(p)
	if err != nil {
		t.Fatal(err)
	}

	if i.Value != 30 {
		t.Errorf("got: <%v>, want: <%v>", i.Value, 30)
	}

	if i.Weight != 9 {
		t.Errorf("got: <%v>, want: <%v>", i.Weight, 9)
	}
}

func TestCrafterGetClass(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	cl, err := c.getClass("art")
	if err != nil {
		t.Fatal(err)
	}

	if cl.Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", cl.Name, "art")
	}
}

func TestCrafterGetClassEmpty(t *testing.T) {
	c := &Crafter{
		classes: map[string]class{},
	}

	_, err := c.getClass("art")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestCrafterGetMaterial(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	m, err := c.getMaterial("wood")
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", m.Name, "wood")
	}
}

func TestCrafterGetMaterialEmpty(t *testing.T) {
	c := &Crafter{
		materials: map[string]material{},
	}

	_, err := c.getMaterial("wood")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestCrafterGetDetail(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	d, err := c.getDetail("beverage")
	if err != nil {
		t.Fatal(err)
	}

	if d.Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", d.Name, "beverage")
	}
}

func TestCrafterGetDetailEmpty(t *testing.T) {
	c := &Crafter{
		details: map[string]detail{},
	}

	_, err := c.getDetail("beverage")
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestCrafterGetDetails(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	d, err := c.getDetails(map[string]string{"contents": "beverage"})
	if err != nil {
		t.Fatal(err)
	}

	if d["contents"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", d["contents"].Name, "beverage")
	}
}

func TestCrafterGetDetailsEmpty(t *testing.T) {
	c := &Crafter{
		details: map[string]detail{},
	}

	_, err := c.getDetails(nil)
	if err == nil {
		t.Error("got: <nil>, want: <error>")
	}
}

func TestCrafterRandomTemplate(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := c.randomTemplate()
	if err != nil {
		t.Fatal(err)
	}

	if tmp.Name == "" {
		t.Errorf("got: <%v>, want: <%v>", "", tmp.Name)
	}
}

func TestCrafterRandomTemplateEmpty(t *testing.T) {
	c := Crafter{
		templates: map[string]template{},
	}

	_, err := c.randomTemplate()
	if err != errEmptyTemplateMap {
		t.Errorf("got: <%v>, want: <%v>", nil, errEmptyTemplateMap)
	}
}

func TestCrafterNew(t *testing.T) {
	tmp, err := os.Open(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}
	defer tmp.Close()

	class, err := os.Open(testFileClass)
	if err != nil {
		t.Fatal(err)
	}
	defer class.Close()

	mat, err := os.Open(testFileMaterial)
	if err != nil {
		t.Fatal(err)
	}
	defer mat.Close()

	qual, err := os.Open(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}
	defer qual.Close()

	det, err := os.Open(testFileDetail)
	if err != nil {
		t.Fatal(err)
	}
	defer det.Close()

	c, err := New(tmp, class, mat, qual, det)
	if err != nil {
		t.Fatal(err)
	}

	if len(c.templates) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(c.templates), 5)
	}

	if c.templates["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", c.templates["figurine"], "figurine")
	}

	if len(c.classes) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(c.classes), 3)
	}

	if c.classes["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", c.classes["art"].Name, "art")
	}

	if len(c.materials) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(c.materials), 8)
	}

	if c.materials["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", c.materials["wood"].Name, "wood")
	}

	if len(c.qualities) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(c.qualities), 7)
	}

	if c.qualities["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", c.qualities["wear"].Name, "wear")
	}

	if len(c.details) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(c.details), 5)
	}

	if c.details["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", c.details["beverage"].Name, "beverage")
	}
}

func TestCrafterNewEmpty(t *testing.T) {
	var emptyFile *os.File

	tmp, err := os.Open(testFileTemplate)
	if err != nil {
		t.Fatal(err)
	}
	defer tmp.Close()

	class, err := os.Open(testFileClass)
	if err != nil {
		t.Fatal(err)
	}
	defer class.Close()

	mat, err := os.Open(testFileMaterial)
	if err != nil {
		t.Fatal(err)
	}
	defer mat.Close()

	qual, err := os.Open(testFileQuality)
	if err != nil {
		t.Fatal(err)
	}
	defer qual.Close()

	det, err := os.Open(testFileDetail)
	if err != nil {
		t.Fatal(err)
	}
	defer det.Close()

	var tests = []struct {
		Name  string
		Tmp   *os.File
		Class *os.File
		Mat   *os.File
		Qual  *os.File
		Det   *os.File
	}{
		{"Empty template reader", emptyFile, class, mat, qual, det},
		{"Empty class reader", tmp, emptyFile, mat, qual, det},
		{"Empty material reader", tmp, class, emptyFile, qual, det},
		{"Empty quality reader", tmp, class, mat, emptyFile, det},
		{"Empty detail reader", tmp, class, mat, qual, emptyFile},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := New(tt.Tmp, tt.Class, tt.Mat, tt.Qual, tt.Det)
			if err == nil {
				t.Error("got: <nil>, want: <error>")
			}

			// Rewind to beginning of files before next run
			tt.Tmp.Seek(0, 0)
			tt.Class.Seek(0, 0)
			tt.Mat.Seek(0, 0)
			tt.Qual.Seek(0, 0)
			tt.Det.Seek(0, 0)
		})
	}
}

func TestCrafterNewFromFiles(t *testing.T) {
	c, err := NewFromFiles(testFileTemplate, testFileClass, testFileMaterial, testFileQuality, testFileDetail)
	if err != nil {
		t.Fatal(err)
	}

	if len(c.templates) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(c.templates), 5)
	}

	if c.templates["figurine"].Name != "figurine" {
		t.Errorf("got: <%v>, want: <%v>", c.templates["figurine"], "figurine")
	}

	if len(c.classes) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(c.classes), 3)
	}

	if c.classes["art"].Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", c.classes["art"].Name, "art")
	}

	if len(c.materials) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(c.materials), 8)
	}

	if c.materials["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", c.materials["wood"].Name, "wood")
	}

	if len(c.qualities) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(c.qualities), 7)
	}

	if c.qualities["wear"].Name != "wear" {
		t.Errorf("got: <%v>, want: <%v>", c.qualities["wear"].Name, "wear")
	}

	if len(c.details) != 5 {
		t.Errorf("got: <%v>, want: <%v>", len(c.details), 5)
	}

	if c.details["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", c.details["beverage"].Name, "beverage")
	}
}

func TestCrafterNewFromFilesEmpty(t *testing.T) {
	fakeFile := "fake file name"
	var tests = []struct {
		Name  string
		Tmp   string
		Class string
		Mat   string
		Qual  string
		Det   string
	}{
		{"Non-existent template file", fakeFile, testFileClass, testFileMaterial, testFileQuality, testFileDetail},
		{"Non-existent class file", testFileTemplate, fakeFile, testFileMaterial, testFileQuality, testFileDetail},
		{"Non-existent material file", testFileTemplate, testFileClass, fakeFile, testFileQuality, testFileDetail},
		{"Non-existent quality file", testFileTemplate, testFileClass, testFileMaterial, fakeFile, testFileDetail},
		{"Non-existent detail file", testFileTemplate, testFileClass, testFileMaterial, testFileQuality, fakeFile},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := NewFromFiles(tt.Tmp, tt.Class, tt.Mat, tt.Qual, tt.Det)
			if err == nil {
				t.Error("got: <nil>, want: <error>")
			}
		})
	}
}
