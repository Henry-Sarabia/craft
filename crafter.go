package craft

import (
	"io"
	"math/rand"

	"github.com/pkg/errors"
)

var errEmptyTemplateMap = errors.New("template map is empty")
var errEmptyDetailMap = errors.New("cannot getDetails from empty ref map")

// Crafter contains all of the available resources to be used when generating
// items.
type Crafter struct {
	templates map[string]template
	classes   map[string]class
	materials map[string]material
	qualities map[string]quality
	details   map[string]detail
}

// NewItem returns a uniquely generated Item using the resources available to
// the Crafter receiver.
func (c *Crafter) NewItem() (*Item, error) {
	tmp, err := c.randomTemplate()
	if err != nil {
		return nil, errors.Wrap(err, "NewItem: cannot select random template from resources")
	}

	i, err := c.generateItem(tmp)
	if err != nil {
		return nil, errors.Wrap(err, "NewItem: cannot generate item")
	}

	return i, nil
}

// generateItem generates a new item out of the provided template.
func (c *Crafter) generateItem(tmp *template) (*Item, error) {
	proto, err := c.toPrototype(tmp)
	if err != nil {
		return nil, err
	}

	i, err := c.toItem(proto)
	if err != nil {
		return nil, err
	}

	err = i.composeDescription()
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (c *Crafter) toPrototype(tmp *template) (*prototype, error) {
	name, err := tmp.randomName()
	if err != nil {
		return nil, err
	}

	cl, err := c.getClass(tmp.Class)
	if err != nil {
		return nil, err
	}

	cf, err := cl.randomConfiguration()
	if err != nil {
		return nil, err
	}

	m, err := tmp.randomMaterial()
	if err != nil {
		return nil, err
	}

	mat, err := c.getMaterial(m)
	if err != nil {
		return nil, err
	}

	d, err := tmp.randomDetails()
	if err != nil {
		return nil, err
	}

	dets, err := c.getDetails(d)
	if err != nil {
		return nil, err
	}

	return &prototype{
		name:     name,
		value:    tmp.BaseValue,
		weight:   tmp.BaseWeight,
		class:    tmp.Class,
		config:   cf,
		material: mat,
		details:  dets,
	}, nil
}

func (c *Crafter) toItem(p *prototype) (*Item, error) {
	i := &Item{}

	var err error

	if i.Material, err = p.getMaterial(); err != nil {
		return nil, err
	}

	if i.Quality, err = p.getQuality(c.qualities); err != nil {
		return nil, err
	}

	if i.Details, err = p.getDetails(); err != nil {
		return nil, err
	}

	if i.Format, err = p.getFormat(); err != nil {
		return nil, err
	}

	if i.Verb, err = p.getVerb(); err != nil {
		return nil, err
	}

	i.Name = p.name
	i.Value = p.getValue()
	i.Weight = p.getWeight()

	return i, nil
}

func (c *Crafter) getClass(name string) (*class, error) {
	cl, ok := c.classes[name]
	if !ok {
		return nil, errors.Errorf("cannot find '%s' in Crafter.classes", name)
	}

	return &cl, nil
}

func (c *Crafter) getMaterial(name string) (*material, error) {
	mat, ok := c.materials[name]
	if !ok {
		return nil, errors.Errorf("cannot find '%s' in Crafter.materials", name)
	}

	return &mat, nil
}

func (c *Crafter) getDetail(name string) (*detail, error) {
	det, ok := c.details[name]
	if !ok {
		return nil, errors.Errorf("cannot find '%s' in Crafter.details", name)
	}

	return &det, nil
}

func (c *Crafter) getDetails(ref map[string]string) (map[string]*detail, error) {
	if len(ref) < 1 {
		return nil, errEmptyDetailMap
	}

	m := make(map[string]*detail)

	for lbl, v := range ref {
		d, err := c.getDetail(v)
		if err != nil {
			return nil, err
		}

		m[lbl] = d
	}

	return m, nil
}

func (c *Crafter) randomTemplate() (*template, error) {
	if len(c.templates) < 1 {
		return nil, errEmptyTemplateMap
	}
	i := rand.Intn(len(c.templates))

	var tmp template
	for _, tmp = range c.templates {
		if i == 0 {
			break
		}
		i--
	}

	return &tmp, nil
}

// New returns a pointer to an initialized Crafter object populated
// with data read from the provided readers.
func New(tmp, class, mat, qual, det io.Reader) (*Crafter, error) {
	var err error
	res := &Crafter{}

	res.templates, err = readTemplates(tmp)
	if err != nil {
		return nil, err
	}

	res.classes, err = readClasses(class)
	if err != nil {
		return nil, err
	}

	res.materials, err = readMaterials(mat)
	if err != nil {
		return nil, err
	}

	res.qualities, err = readQualities(qual)
	if err != nil {
		return nil, err
	}

	res.details, err = readDetails(det)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// NewFromFiles returns a pointer to an initialized Crafter object populated
// with data found using the provided file names.
func NewFromFiles(tmpFile, classFile, matFile, qualFile, detFile string) (*Crafter, error) {
	tmp, err := loadTemplates(tmpFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load item templates from file named '%s'", tmpFile)
	}

	class, err := loadClasses(classFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load item classes from file named '%s'", classFile)
	}

	mat, err := loadMaterials(matFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load materials from file named '%s'", matFile)
	}

	qual, err := loadQualities(qualFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load qualities from file named '%s'", qualFile)
	}

	det, err := loadDetails(detFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load details from file named '%s'", detFile)
	}

	return &Crafter{
		templates: tmp,
		classes:   class,
		materials: mat,
		qualities: qual,
		details:   det,
	}, nil
}
