package craft

import (
	"io"
	"math/rand"

	"github.com/pkg/errors"
)

var errEmptyTemplateMap = errors.New("itemTemplate map is empty")

// Resources contains all of the available resources to be used when generating
// items.
type Resources struct {
	itemTemplates map[string]itemTemplate
	itemClasses   map[string]itemClass
	qualities     map[string]quality
	materials     map[string]material
	details       map[string]detail
}

func (r *Resources) NewItem() (*Item, error) {
	tmp, err := r.selectTemplate()
	if err != nil {
		return nil, errors.Wrap(err, "cannot select a template from resources")
	}

	i, err := generateItem(tmp, r)
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate item")
	}

	return i, nil
}

func (r *Resources) selectTemplate() (*itemTemplate, error) {
	if len(r.itemTemplates) < 1 {
		return nil, errEmptyTemplateMap
	}
	i := rand.Intn(len(r.itemTemplates))

	// fix this. why set itemtemplate ever iteration? why decrementing?
	var tmp itemTemplate
	for _, tmp = range r.itemTemplates {
		if i == 0 {
			break
		}
		i--
	}

	return &tmp, nil
}

// ReadResources returns a pointer to an initialized Resources object populated
// with data read from the provided readers.
func ReadResources(temp, class, mat, det, qual io.Reader) (*Resources, error) {
	var err error
	rp := &Resources{}

	rp.itemTemplates, err = readItemTemplates(temp)
	if err != nil {
		return nil, err
	}

	rp.itemClasses, err = readItemClasses(class)
	if err != nil {
		return nil, err
	}

	rp.materials, err = readMaterials(mat)
	if err != nil {
		return nil, err
	}

	rp.details, err = readDetails(det)
	if err != nil {
		return nil, err
	}

	rp.qualities, err = readQualities(qual)
	if err != nil {
		return nil, err
	}

	return rp, nil
}

// LoadResources returns a pointer to an initialized Resources object populated
// with data found using the provided file names.
func LoadResources(tempFile, classFile, matFile, detFile, qualFile string) (*Resources, error) {
	temp, err := loadItemTemplates(tempFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load item templates from file named '%s'", tempFile)
	}

	class, err := loadItemClasses(classFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load item classes from file named '%s'", classFile)
	}

	mat, err := loadMaterials(matFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load materials from file named '%s'", matFile)
	}

	det, err := loadDetails(detFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load details from file named '%s'", detFile)
	}

	qual, err := loadQualities(qualFile)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load qualities from file named '%s'", qualFile)
	}

	return &Resources{
		itemTemplates: temp,
		itemClasses:   class,
		materials:     mat,
		details:       det,
		qualities:     qual,
	}, nil
}
